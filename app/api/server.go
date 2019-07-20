package server

import (
	"github.com/andresoro/rlog/app/store"
	"github.com/gorilla/mux"
)

// Server is the main entry point for our api
type Server struct {
	api *mux.Router
	db  *store.Store
}
