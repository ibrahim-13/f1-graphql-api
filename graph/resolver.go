//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"f1-gql-api/f1api"
	"f1-gql-api/graph/model"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Api *f1api.F1Api
}

func NewResolver() *Resolver {
	return &Resolver{Api: f1api.NewF1Api()}
}

func (resolver *Resolver) GetRaces(filter *model.RaceFilter) ([]*model.Race, error) {
	r, err := resolver.Api.GetRaceListByYear("2023")
	if err != nil {
		return nil, err
	}
	var races []*model.Race
	for i := range r {
		races = append(races, mapToRace(r[i]))
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
