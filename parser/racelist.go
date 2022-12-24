package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

const (
	__url_race_list string = "https://www.formula1.com/en/racing/%s.html"
	__time_layout_1 string = "2006-01-02T15:04:05"
)

const ()

type RaceLocation struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Race struct {
	Url           string       `json:"@id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	StartDate     string       `json:"startDate"`
	StartDateTime time.Time    `json:"-"`
	EndDate       string       `json:"endDate"`
	EndDateTime   time.Time    `json:"-"`
	Location      RaceLocation `json:"location"`
}

type RaceEvent struct {
	Url           string    `json:"@id"`
	Name          string    `json:"name"`
	StartDate     string    `json:"startDate"`
	StartDateTime time.Time `json:"-"`
	EndDate       string    `json:"endDate"`
	EndDateTime   time.Time `json:"-"`
}

type RaceEventData struct {
	Race
	SubEvents []RaceEvent `json:"subEvent"`
}

func GetRaceList(year string) ([]Race, error) {
	r, e := getLinkedData[Race](fmt.Sprintf(__url_race_list, year))
	if e != nil {
		return nil, e
	}
	for i := range r {
		r[i].StartDateTime, _ = time.Parse(__time_layout_1, r[i].StartDate)
		r[i].EndDateTime, _ = time.Parse(__time_layout_1, r[i].EndDate)
	}
	return r, nil
}

func (race Race) GetRaceEventList() ([]RaceEventData, error) {
	r, e := getLinkedData[RaceEventData](race.Url)
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

func getLinkedData[K interface{}](url string) ([]K, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	var races []K
	stack := &NodeStack{}
	stack.Push(doc)
	isHeadFound := false
	for {
		if stack.IsEmpty() {
			break
		}
		current, _ := stack.Pop()
		if isLinkedDataTag(current) {
			var race K
			err = json.Unmarshal([]byte(current.FirstChild.Data), &race)
			if err != nil {
				panic(err)
			}
			races = append(races, race)
		}
		if !isHeadFound {
			for c := current.FirstChild; c != nil; c = c.NextSibling {
				stack.Push(c)
			}
		}
		isHeadFound = isHeadTag(current)
	}
	return races, nil
}

func isLinkedDataTag(node *html.Node) bool {
	if node != nil && node.Type == html.ElementNode && node.Data == "script" {
		for i := range node.Attr {
			if node.Attr[i].Key == "type" &&
				node.Attr[i].Val == "application/ld+json" &&
				node.FirstChild != nil &&
				node.FirstChild.Data != "" {
				return true
			}
		}
	}
	return false
}

func isHeadTag(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "head"
}
