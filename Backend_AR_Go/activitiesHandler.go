package main

import (
	"encoding/json"
	"net/http"
)

func getActivitiesHandler(w http.ResponseWriter, r *http.Request) {

	checkErr := checkWeeklyReset()
	if checkErr != nil {
		http.Error(w, checkErr.Error(), http.StatusInternalServerError)
		return
	}

	activities, err := readActivities("activitiesLog.json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp = map[string]interface{}{
		"activities": activities,
	}
	json.NewEncoder(w).Encode(resp)
}
