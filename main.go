package main

import (
	"log"
	"net/http"

	"github.com/andresoro/rlog/app/api"
	"github.com/andresoro/rlog/app/store/postgres"
)

func main() {
	db := &postgres.DB{}

	srv := api.New(db)
	err := srv.Init()
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", srv.Mux))
}
