package main

import (
	"fmt"
	"time"

	"f1-gql-api/f1api"
)

func PrintEvents() {
	api := f1api.NewF1Api()
	races, err := api.GetRaceList("2023")
	if err != nil {
		panic(err)
	} else {
		if len(races) > 0 {
			event, _ := api.GetRaceEventList(races[0])
			fmt.Println("Start Time", races[0].StartDateTime.Format(time.RFC3339))
			fmt.Println("Events", len(event))
		}
	}
}
