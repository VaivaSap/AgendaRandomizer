package main

import (
	"fmt"
	"math/rand"
)

func randomizeActivities(activities []Activity) Activity {
	var totalWeight float64 = 0

	for _, activity := range activities {
		remainingHours := activity.TargetHours - activity.LoggedHours
		if remainingHours <= 0 {
			remainingHours = 0
		}
		totalWeight = totalWeight + remainingHours
	}

	var random = rand.Float64() * totalWeight
	fmt.Println("Random point:", random)

	var selectedActivity Activity

	for _, activity := range activities {
		remainingHours := activity.TargetHours - activity.LoggedHours

		if remainingHours <= 0 { 
			remainingHours = 0 
		}

		if random-remainingHours > 0 {
			random = random - remainingHours
		} else {
			selectedActivity = activity
			break
		}
	}

	return selectedActivity
}
