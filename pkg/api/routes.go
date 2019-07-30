package api

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Routes initializes and returns api router
func (a *API) Routes() {

	// init
	r := mux.NewRouter()

	// singular event methods
	r.HandleFunc("/events/{id}", a.GetEvent).Methods("GET")
	r.HandleFunc("/sites/{id}/events", a.AllEvents).Methods("GET")

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
		return
	}

	// write json back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// AllEvents returns all site iteractions for a site with given id
// and a given time frame
func (a *API) AllEvents(w http.ResponseWriter, r *http.Request) {

	// site id parameter
	ids := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// read request body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// model req body as this struct
	var req struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// hit db for event aggregates
	agg, err := a.db.RetrieveSiteStats(req.Start, req.End, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set headers and write json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(agg)
}
