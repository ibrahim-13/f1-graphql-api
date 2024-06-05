//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"f1-gql-api/f1api"
	"f1-gql-api/graph/model"
	"fmt"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Api *f1api.F1Api
}

func NewResolver() *Resolver {
	year, cache := fmt.Sprint(time.Now().Year()), f1api.NewApiCacheLocal(6*time.Hour)
	return &Resolver{Api: f1api.NewF1Api(year, cache)}
}

func (resolver *Resolver) GetRaces(filter *model.RaceFilter) ([]*model.Race, error) {
	r, err := resolver.Api.GetRaceList()
	if err != nil {
		return nil, err
	}
	var races []*model.Race
	currentTime := time.Now()
	switch *filter {
	case model.RaceFilterAllNextRace:
		for i := range r {
			if r[i].StartDateTime.After(currentTime) {
				races = append(races, mapToRace(r[i]))
			}
		}
	case model.RaceFilterAllRace:
		for i := range r {
			races = append(races, mapToRace(r[i]))
		}
	case model.RaceFilterOnlyNextRace:
		for i := range r {
			if r[i].StartDateTime.After(currentTime) {
				races = append(races, mapToRace(r[i]))
				break
			}
		}
	}
	return races, nil
}

func (resolver *Resolver) GetRaceEvents(race *model.Race) ([]*model.RaceEvent, error) {
	r, err := resolver.Api.GetRaceEventList(race.URL)
	if err != nil {
		return nil, err
	}
	var events []*model.RaceEvent
	if len(r) < 1 {
		return nil, nil
	}
	for i := range r[0].SubEvents {
		events = append(events, &model.RaceEvent{
			URL:   r[0].SubEvents[i].Url,
			Name:  r[0].SubEvents[i].Name,
			Start: r[0].SubEvents[i].StartDateTime.Format(time.RFC3339),
			End:   r[0].SubEvents[i].EndDateTime.Format(time.RFC3339),
		})
	}
	return events, nil
}
