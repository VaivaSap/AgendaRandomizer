package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hourLoggingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var logRequest LogEntry
	var err error

	err = json.NewDecoder(r.Body).Decode(&logRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if logRequest.HoursSpent <= 0 || logRequest.ActivityName == "" {
		http.Error(w, "Invalid log entry", http.StatusBadRequest)
		return
	}

	fmt.Println(logRequest)

	activities, err := readActivities("activitiesLog.json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(activities)
	activityFound := false
	for index, value := range activities {
		fmt.Println(index, value)

		if value.ActivityName == logRequest.ActivityName {
			activities[index].LoggedHours += logRequest.HoursSpent
			activityFound = true
			fmt.Println("Activity found", activities[index].LoggedHours)
			break
		}
	}

	if !activityFound {
		http.Error(w, "Activity not found", http.StatusNotFound)
		return
	}

	err = writeActivities("activitiesLog.json", activities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Activity updated successfully"})
}
