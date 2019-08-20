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

func (a *API) GetSites(w http.ResponseWriter, r *http.Request) {

}
