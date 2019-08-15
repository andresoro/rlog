// Package model holds definitions for our api and database
package model

import "time"

// Site hold meta information about websites being tracked
type Site struct {
	ID   int
	Name string
}

// Event is an individual request made to a particular website
// holds information relevant about the given session
type Event struct {
	ID     int     `json:"id"`
	SiteID int     `json:"site_id"`
	Key    string    `json:"key"`
	Addr   string    `json:"addr"`
	Date   time.Time `json:"date"`
	Unique bool      `json:"unique"`
}

// SiteStats represent aggregated stats of a website for a given timeframe
type SiteStats struct {
	SiteID int
	Pages  []*KeyStats
}

// KeyStats are aggregated statistics for a given path and a given timeframe
type KeyStats struct {
	SiteID       int
	Key          string
	UniqueVisits int
	PageViews    int
	Start        time.Time
	End          time.Time
}

// Add an event to pagestat aggregate
func (ps *KeyStats) Add(e *Event) {
	if e.Unique {
		ps.UniqueVisits++
	}

	ps.PageViews++

}
