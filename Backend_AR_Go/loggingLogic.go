package main

import (
	"fmt"
	"os"
	"encoding/json"
)

func readActivities(filePath string) ([]Activity, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
        return nil, err
	}

	var activities []Activity

	err = json.Unmarshal(data, &activities)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	fmt.Println(activities)
	return activities, nil
}

func writeActivities(filePath string, activities []Activity) error {
	data, err := json.Marshal(activities)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}