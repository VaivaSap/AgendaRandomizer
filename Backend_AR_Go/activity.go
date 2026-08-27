package main

type Activity struct {
	ActivityName string  `json:"activity_name"`
	TargetHours  float64 `json:"target_hours"`
	LoggedHours  float64 `json:"logged_hours"`
}
