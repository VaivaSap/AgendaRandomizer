package main

import (
	"fmt"
	"math/rand"
)

func randomizeActivities(activities []Activity) Activity {
	var totalWeight float64 = 0

	for _, activity := range activities {
		totalWeight = totalWeight + activity.Hours
	}

	var random = rand.Float64() * totalWeight
	fmt.Println("Random point:", random)

	var selectedActivity Activity

	for _, activity := range activities {
		if random-activity.Hours > 0 {
			random = random - activity.Hours
		} else {
			selectedActivity = activity
			break
		}
	}

	return selectedActivity
}
