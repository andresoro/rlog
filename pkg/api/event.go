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

// GetEvent is the GET endpoint that returns a specific event based on
// the ID in the URL
func (a *API) GetEvent(w http.ResponseWriter, r *http.Request) {
	// siteID and event id parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// convert id param to int
	eventID, err := strconv.Atoi(id)
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// write json back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// DatedEvents returns all site iteractions for a site with given id
// and a given time frame
func (a *API) DatedEvents(w http.ResponseWriter, r *http.Request) {

	// site id parameter
	ids := mux.Vars(r)["id"]
	id, err := strconv.Atoi(ids)
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
	// if our body contains no start/end date, return events in past 24 hours
	if len(b) > 0 {
		err = json.Unmarshal(b, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		req.Start = time.Now().AddDate(0, 0, -1)
		req.End = time.Now()
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

// AllEvents returns all events for a siteID
func (a *API) AllEvents(w http.ResponseWriter, r *http.Request) {

	// site id parameter
	ids := mux.Vars(r)["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("input a correct site ID"))
		return
	}

	events, err := a.db.RetrieveAll(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
