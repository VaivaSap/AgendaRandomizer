package main

import (
	"net/http"
)

func main() {
	http.Handle("/activities", corsMiddleware(http.HandlerFunc(getActivitiesHandler)))
	http.Handle("/randomize", corsMiddleware(http.HandlerFunc(randomizeActivitiesHandler)))
	http.ListenAndServe(":8080", nil)
}
