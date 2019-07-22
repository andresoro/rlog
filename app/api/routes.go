package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes initializes and returns api router
func (a *API) Routes() {

	// init
	r := mux.NewRouter()

	r.HandleFunc("/site/{id}", a.GetSite).Methods(http.MethodGet)
	r.HandleFunc("/site/", a.PostSite).Methods(http.MethodPost)

	r.HandleFunc("/site/{id}/stats", a.GetSiteStats).Methods(http.MethodGet)

	a.Mux = r
}

// GetSite /site/{id} returns information about site with given ID
func (a *API) GetSite(w http.ResponseWriter, r *http.Request) {

}

// PostSite /site/ recieves data about a newly registered site
func (a *API) PostSite(w http.ResponseWriter, r *http.Request) {

}

// GetSiteStats /site/{id}/stats returns data about website traffic within a timeframe
// timeframe can either be passed or default
func (a *API) GetSiteStats(w http.ResponseWriter, r *http.Request) {

}
