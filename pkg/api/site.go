package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// NewSite endpoint will register a new site
func (a *API) NewSite(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	id, err := a.db.InsertSite(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

// GetSites handler will return an array of site metadata (ID, name)
func (a *API) GetSites(w http.ResponseWriter, r *http.Request) {
	sites, err := a.db.GetSites()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sites)
}
