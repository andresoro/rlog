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
	RetrieveEvent(id int) (*model.Event, error)
	RetrieveSiteStats(start, end time.Time, siteID int) (*model.SiteStats, error)
	RetrieveAll(siteID int) ([]*model.Event, error)
	InsertSite(name string) (int, error)
}
