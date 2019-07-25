package store

import (
	model "github.com/andresoro/rlog/app/models"
)

// Store - interface for handling app models against an arbitrary database
// this interface is, ideally, the contract between the api and the database
type Store interface {

	// DB model methods to actually write and retrieve our data
	InsertEvent(*model.Event) error
	RetrieveEvent(id int64) (*model.Event, error)
}
