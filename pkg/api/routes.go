package api

import (
	"github.com/gorilla/mux"
)

// Routes initializes and returns api router
func (a *API) Routes() {

	// init
	r := mux.NewRouter()

	// methods to return event data from database
	r.HandleFunc("/events/{id}", a.GetEvent).Methods("GET")
	r.HandleFunc("/sites/{id}", a.AllEvents).Methods("GET")

	// collection endpoint to recieve tracking data from clients
	r.PathPrefix("/collect").HandlerFunc(a.Collect)

	a.Mux = r
}
