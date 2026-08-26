package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/activities", getActivitiesHandler)
	http.HandleFunc("/randomize", randomizeActivitiesHandler)
    http.ListenAndServe(":8080", nil)
}


