package f1api

import (
	"context"
	"fmt"
	"time"
)

const (
	__time_layout_1 string = "2006-01-02T15:04:05"
)

type F1Api struct {
	year              string
	cacheDuration     time.Duration
	apiRequestTimeout time.Duration
}

func NewF1Api() *F1Api {
	return &F1Api{
		year:              fmt.Sprint(time.Now().Year()),
		cacheDuration:     24 * time.Hour,
		apiRequestTimeout: 8 * time.Second,
	}
}

func (ctx *F1Api) GetRaceList(year string) ([]Race, error) {
	return ctx.GetRaceListByYear(ctx.year)
}

func (f1 *F1Api) GetRaceListByYear(year string) ([]Race, error) {
	ctx, _ := f1.getRequestCtx()
	r, e := getLinkedData[Race](ctx, fmt.Sprintf(__url_race_list, year))
	if e != nil {
		return nil, e
	}
	for i := range r {
		r[i].StartDateTime, _ = time.Parse(__time_layout_1, r[i].StartDate)
		r[i].EndDateTime, _ = time.Parse(__time_layout_1, r[i].EndDate)
	}
	return r, nil
}

func (f1 *F1Api) GetRaceEventList(race Race) ([]RaceEventData, error) {
	ctx, _ := f1.getRequestCtx()
	r, e := getLinkedData[RaceEventData](ctx, race.Url)
	if e != nil {
		return nil, e
	}
	for i := range r {
		r[i].StartDateTime, _ = time.Parse(time.RFC3339, r[i].StartDate)
		r[i].EndDateTime, _ = time.Parse(time.RFC3339, r[i].EndDate)
		for j := range r[i].SubEvents {
			r[i].SubEvents[j].StartDateTime, _ = time.Parse(time.RFC3339, r[i].SubEvents[j].StartDate)
			r[i].SubEvents[j].EndDateTime, _ = time.Parse(time.RFC3339, r[i].SubEvents[j].EndDate)
		}
	}

	return r, nil
}

func (f1 *F1Api) getRequestCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), f1.apiRequestTimeout)
}
