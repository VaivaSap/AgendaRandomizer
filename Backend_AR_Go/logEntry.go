package main

import "time"

type LogEntry struct {
	ActivityName string    `json:"activity_name"`
	HoursSpent   float64   `json:"hours_spent"`
	Completed    bool      `json:"completed"`
	Date         time.Time `json:"date"`
}
