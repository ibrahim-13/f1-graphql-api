package f1api

import (
	"errors"
	"time"
)

var ErrCacheRetentionTimeExceeded = errors.New("err_cache_retention_time_exccede")
var ErrCacheNotFount = errors.New("err_cache_not_found")

type cacheData[K interface{}] struct {
	time time.Time
	data []K
}

type ResponseCache struct {
	retentionTime  time.Duration
	cacheRace      map[string]cacheData[Race]
	cacheRaceEvent map[string]cacheData[RaceEventData]
}

func NewResponseCache(retentionTime time.Duration) *ResponseCache {
	return &ResponseCache{
		retentionTime:  retentionTime,
		cacheRace:      map[string]cacheData[Race]{},
		cacheRaceEvent: map[string]cacheData[RaceEventData]{},
	}
}

func (c *ResponseCache) SetRace(url string, race []Race) {
	c.cacheRace[url] = cacheData[Race]{time: time.Now(), data: race}
}

func (c *ResponseCache) GetRace(url string) ([]Race, error) {
	data, ok := c.cacheRace[url]
	if !ok {
		return nil, ErrCacheNotFount
	}
	if data.time.Before(time.Now()) {
		return nil, ErrCacheRetentionTimeExceeded
	}
	return data.data, nil
}

func (c *ResponseCache) SetRaceEvent(url string, event []RaceEventData) {
	c.cacheRaceEvent[url] = cacheData[RaceEventData]{time: time.Now(), data: event}
}

func (c *ResponseCache) GetRaceEvent(url string) ([]RaceEventData, error) {
	data, ok := c.cacheRaceEvent[url]
	if !ok {
		return nil, ErrCacheNotFount
	}
	if data.time.Before(time.Now()) {
		return nil, ErrCacheRetentionTimeExceeded
	}
	return data.data, nil
}
