package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andresoro/rlog/pkg/api"
	"github.com/andresoro/rlog/pkg/store/postgres"
)

func main() {
	// load postgres with envirornment variables
	db := &postgres.DB{
		Config: &postgres.Config{
			User:      os.Getenv("DB_USER"),
			Pass:      os.Getenv("DB_PASS"),
			DB:        os.Getenv("DB_NAME"),
			Migration: os.Getenv("MIGRATION_DIR"),
		},
	}

	// ensure connection to postgres
	err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %e", err)
	}

	// apply any database migrations that may have been introduced
	err = db.Migrate()
	if err != nil {
		log.Fatalf("Could not perform database migrations: %e", err)
	}

	// init server
	server := api.New(db)

	// run
	log.Fatal(http.ListenAndServe(":8080", server.Mux))
}
