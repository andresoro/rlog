package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Routes initializes and returns api router
func (a *API) Routes() {

	// init
	r := mux.NewRouter()

	// methods to return event data from database
	r.HandleFunc("/events/{id}", a.GetEvent).Methods("GET")
	r.HandleFunc("/sites/{id}", a.DatedEvents).Methods("GET")
	r.HandleFunc("/sites/all/{id}", a.AllEvents).Methods("GET")
	r.HandleFunc("/sites/new/{name}", a.NewSite).Methods("GET")

	// collection catch all endpoint to recieve tracking data from clients
	// this endpoint serves a 1x1 pixel
	r.PathPrefix("/collect").HandlerFunc(a.Collect)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(os.Getenv("FRONTEND_DIR"))))

	a.Mux = r
}
