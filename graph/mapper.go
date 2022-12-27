package graph

import (
	"f1-gql-api/f1api"
	"f1-gql-api/graph/model"
	"time"
)

func mapToRace(race f1api.Race) *model.Race {
	return &model.Race{
		URL:         race.Url,
		Name:        race.Name,
		Description: race.Description,
		Start:       race.StartDateTime.Format(time.RFC3339),
		End:         race.EndDateTime.Format(time.RFC3339),
	}
}
