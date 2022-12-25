// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Race struct {
	URL        string       `json:"url"`
	Name       string       `json:"name"`
	Descriptin string       `json:"descriptin"`
	Start      string       `json:"start"`
	End        string       `json:"end"`
	Events     []*RaceEvent `json:"events"`
}

type RaceEvent struct {
	URL   string `json:"url"`
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
}
