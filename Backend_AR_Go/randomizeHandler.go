package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func randomizeActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	activities, err := readActivities("activitiesLog.json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := randomizeActivities(activities)
	fmt.Println("Do this:", result)

	json.NewEncoder(w).Encode(result)
}
