package main

import (
	"fmt"
)

func main() {
	activities := []Activity{
		{ActivityName: "A", Hours: 45.0},
		{ActivityName: "B", Hours: 15.0},
		{ActivityName: "C", Hours: 10.0},
		{ActivityName: "D", Hours: 4.0},
		{ActivityName: "E", Hours: 3.0},
		{ActivityName: "F", Hours: 3.0},
		{ActivityName: "G", Hours: 10.0},
		{ActivityName: "H", Hours: 7.0},
		{ActivityName: "I", Hours: 5.0},
		{ActivityName: "J", Hours: 5.0},
	}

	result := randomizeActivities(activities)
	fmt.Println("Do this:", result)
}
