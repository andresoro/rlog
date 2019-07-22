package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andresoro/rlog/app/api"
	"github.com/andresoro/rlog/app/store/postgres"
)

func main() {
	// load postgres with needed variables
	db := &postgres.DB{
		Config: &postgres.Config{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			DB:   os.Getenv("DB_NAME"),
		},
	}

	srv := api.New(db)
	err := srv.Init()
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", srv.Mux))
}
