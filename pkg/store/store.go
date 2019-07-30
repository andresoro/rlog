package store

import (
	"time"

	model "github.com/andresoro/rlog/pkg/model"
)

// Store - interface for handling app model against an arbitrary database
// this interface is, ideally, the contract between the api and the database
type Store interface {

	// DB model methods to actually write and retrieve our data
	InsertEvent(*model.Event) error
	RetrieveEvent(id int64) (*model.Event, error)
	RetrieveSiteStats(start, end time.Time, siteID int64) (*model.SiteStats, error)
}
