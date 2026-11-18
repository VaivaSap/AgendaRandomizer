package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func weeklyTrackingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	loc, err := time.LoadLocation("Europe/Vilnius")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	now := time.Now().In(loc)
	weekStart := now.Format("2006-01-02")

	trackingLog := TrackingLog{WeekStart: weekStart}

	data, err := json.Marshal(trackingLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("weekStart.json", data, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "week_start": weekStart})
}

func checkWeeklyReset() error {
	data, err := os.ReadFile("weekStart.json")

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var trackingLog TrackingLog
	err = json.Unmarshal(data, &trackingLog)
	if err != nil {
		return err
	}

	start, err := time.Parse("2006-01-02", trackingLog.WeekStart)
	if err != nil {
		return err
	}

	loc, err := time.LoadLocation("Europe/Vilnius")
	if err != nil {
		return err
	}
	todayText := time.Now().In(loc).Format("2006-01-02")
	today, err := time.Parse("2006-01-02", todayText)
	if err != nil {
		return err
	}

	days := int(today.Sub(start).Hours() / 24)
	if days < 7 {
		return nil
	}
	err = resetHours()
	if err != nil {
		return err
	}

	weeks := days / 7
	newStart := start.AddDate(0, 0, weeks*7)

	newLog := TrackingLog{WeekStart: newStart.Format("2006-01-02")}
	data, err = json.Marshal(newLog)
	if err != nil {
		return err
	}
	err = os.WriteFile("weekStart.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func resetHours() error {
	activities, err := readActivities("activitiesLog.json")
	if err != nil {
		fmt.Println("Resetting failed", err)
		return err
	}
	for i := range activities {
		activities[i].LoggedHours = 0
	}
	return writeActivities("activitiesLog.json", activities)
}
