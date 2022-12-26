package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"f1-gql-api/graph/model"
	"log"
)

// Races is the resolver for the races field.
func (r *queryResolver) Races(ctx context.Context) ([]*model.Race, error) {
	log.Println("Resolver : Race")
	return []*model.Race{{URL: "url", Name: ""}}, nil
}

// Events is the resolver for the events field.
func (r *raceResolver) Events(ctx context.Context, obj *model.Race) ([]*model.RaceEvent, error) {
	log.Println("Resolver : RaceEvent")
	return []*model.RaceEvent{{URL: "url"}}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Race returns RaceResolver implementation.
func (r *Resolver) Race() RaceResolver { return &raceResolver{r} }

type queryResolver struct{ *Resolver }
type raceResolver struct{ *Resolver }
