package main

import (
	"fmt"
	"time"

	"f1api/parser"
)

func PrintEvents() {
	races, err := parser.GetRaceList("2023")
	if err != nil {
		panic(err)
	} else {
		if len(races) > 0 {
			event, _ := races[0].GetRaceEventList()
			fmt.Println("Start Time", races[0].StartDateTime.Format(time.RFC3339))
			fmt.Println("Events", len(event))
		}
	}
}
