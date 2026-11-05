package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/activities", corsMiddleware(http.HandlerFunc(getActivitiesHandler)))
	http.Handle("/randomize", corsMiddleware(http.HandlerFunc(randomizeActivitiesHandler)))
	http.Handle("/log", corsMiddleware(http.HandlerFunc(hourLoggingHandler)))
	
	fmt.Println("Server listening on :8080")
	var err error = http.ListenAndServe(":8080", nil)
	fmt.Println("Server stopped:", err)
}
