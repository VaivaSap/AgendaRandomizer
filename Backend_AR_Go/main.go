package main

import (
	"fmt"
	"net/http"
)

func main() {
checkErr := checkWeeklyReset()
	if checkErr != nil {
		fmt.Println("Weekly check failed:", checkErr)
	}

	http.Handle("/activities", corsMiddleware(http.HandlerFunc(getActivitiesHandler)))
	http.Handle("/randomize", corsMiddleware(http.HandlerFunc(randomizeActivitiesHandler)))
	http.Handle("/log", corsMiddleware(http.HandlerFunc(hourLoggingHandler)))
	http.Handle("/weeklyTracking", corsMiddleware(http.HandlerFunc(weeklyTrackingHandler)))

	fmt.Println("Server listening on :8080")
	var err error = http.ListenAndServe(":8080", nil)
	fmt.Println("Server stopped:", err)
}
