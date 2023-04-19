package f1api

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	__time_layout_1 string = "2006-01-02T15:04:05"
)

type F1Api struct {
	year              string
	apiRequestTimeout time.Duration
	cache             *ResponseCache
}

func NewF1Api() *F1Api {
	return &F1Api{
		year:              fmt.Sprint(time.Now().Year()),
		apiRequestTimeout: 8 * time.Second,
		cache:             NewResponseCache(6 * time.Hour),
	}
}

func (ctx *F1Api) GetRaceList(year string) ([]Race, error) {
	return ctx.GetRaceListByYear(ctx.year)
}

func (f1 *F1Api) GetRaceListByYear(year string) ([]Race, error) {
	url := fmt.Sprintf(__url_race_list, year)
	cached, err := f1.cache.GetRace(url)
	if err == nil {
		return cached, nil
	}
	ctx, _ := f1.getRequestCtx()
	r, e := getLinkedData[Race](ctx, url)
	if e != nil {
		return nil, e
	}
	for i := range r {
		r[i].StartDateTime, _ = time.Parse(__time_layout_1, r[i].StartDate)
		r[i].EndDateTime, _ = time.Parse(__time_layout_1, r[i].EndDate)
		r[i].UpdateNameAndDescription()
	}
	f1.cache.SetRace(url, r)
	return r, nil
}

func (f1 *F1Api) GetRaceEventList(raceUrl string) ([]RaceEventData, error) {
	cached, err := f1.cache.GetRaceEvent(raceUrl)
	if err == nil {
		return cached, nil
	}
	ctx, _ := f1.getRequestCtx()
	r, e := getLinkedData[RaceEventData](ctx, raceUrl)
	if e != nil {
		return nil, e
	}
	for i := range r {
		r[i].StartDateTime, _ = time.Parse(time.RFC3339, r[i].StartDate)
		r[i].EndDateTime, _ = time.Parse(time.RFC3339, r[i].EndDate)
		for j := range r[i].SubEvents {
			r[i].SubEvents[j].StartDateTime, _ = time.Parse(time.RFC3339, r[i].SubEvents[j].StartDate)
			r[i].SubEvents[j].EndDateTime, _ = time.Parse(time.RFC3339, r[i].SubEvents[j].EndDate)
			r[i].SubEvents[j].Name = f1.getSubEventName(r[i].SubEvents[j].Name)
		}
	}
	f1.cache.SetRaceEvent(raceUrl, r)
	return r, nil
}

func (f1 *F1Api) getSubEventName(name string) string {
	s := strings.Split(name, "-")
	if len(s) > 1 {
		return strings.TrimSpace(s[0])
	}
	return name
}

func (f1 *F1Api) getRequestCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), f1.apiRequestTimeout)
}
