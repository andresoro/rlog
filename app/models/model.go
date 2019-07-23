// Package model holds definitions for our api and database
package models

import "time"

// Site hold meta information about websites being tracked
type Site struct {
	ID   int64
	Name string
}

// SiteStats represent aggregated stats of a website for a given timeframe
type SiteStats struct {
	SiteID int64
	Pages  []*PageStats
}

// PageStats are aggregated statistics for a given path and a given timeframe
type PageStats struct {
	SiteID       int64
	Host         string
	Path         string
	UniqueVisits int64
	PageViews    int64
	Start        time.Time
	End          time.Time
}

// PageView is an individual request made to a particular website
// holds information relevant about given sessions
type PageView struct {
	SiteID   int64
	StatID   int64
	Duration int64
	Host     string
	Path     string
	Referrer string
	Unique   bool
	Date     time.Time
}
