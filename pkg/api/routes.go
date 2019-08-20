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

	// handle all api paths under this router
	api := r.PathPrefix("/api").Subrouter()

	// methods to return event data from database
	api.HandleFunc("/events/{id}", a.GetEvent).Methods("GET")
	api.HandleFunc("/sites/{id}", a.DatedEvents).Methods("GET")
	api.HandleFunc("/sites/all/{id}", a.AllEvents).Methods("GET")
	api.HandleFunc("/sites/new/{name}", a.NewSite).Methods("GET")
	api.HandleFunc("/sites/all", a.GetSites).Methods("GET")

	// collection catch all endpoint to recieve tracking data from clients
	// this endpoint serves a 1x1 pixel
	api.PathPrefix("/collect").HandlerFunc(a.Collect)

	// handle frontend static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(os.Getenv("FRONTEND_DIR"))))

	a.Mux = r
}
