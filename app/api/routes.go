package api

import "github.com/gorilla/mux"

// Routes initializes and returns api router
func (a *API) Routes() {
	r := mux.NewRouter()
	a.Mux = r
}
