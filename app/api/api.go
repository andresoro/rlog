package api

import (
	"github.com/andresoro/rlog/app/store"
	"github.com/gorilla/mux"
)

// API represents our server instance that routes requests and returns db data
type API struct {
	db  store.Store
	Mux *mux.Router
}

// New returns a new API instance with access to a db interface
func New(db store.Store) *API {
	api := &API{
		db: db,
	}

	return api

}

// Init will open the DB connection and initialize routes
func (api *API) Init() error {
	err := api.db.Connect()
	if err != nil {
		return err
	}

	api.Routes()

	return nil
}
