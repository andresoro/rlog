package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Routes initializes and returns api router
func (a *API) Routes() {

	// init
	r := mux.NewRouter()

	// singular event methods
	r.HandleFunc("/events/{id}", a.GetEvent).Methods("GET")

	a.Mux = r
}

// GetEvent is the GET endpoint that returns a specific event based on
// the ID in the URL
func (a *API) GetEvent(w http.ResponseWriter, r *http.Request) {
	// siteID and event id parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// convert id param to int64
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make db call
	event, err := a.db.RetrieveEvent(eventID)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// write json back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}
