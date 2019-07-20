package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

// Connect to postgres db, uses ENV variables from docker image
func (db *DB) Connect() error {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB")

	s := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, user, pass, name)
	conn, err := sql.Open("postgres", s)
	if err != nil {
		return err
	}

	db.conn = conn

	log.Println("Attempting to register connection with postgres db...")
	for i := 0; i < 5; i++ {
		err := db.Ping()
		// if connection is working, return
		if err == nil {
			log.Println("Successfully connected to postgres db!")
			return nil
		}
		// wait one second till retrying
		log.Println("Ping to db failed, sleeping one second and trying again")
		time.Sleep(time.Second)
	}

	return errors.New("Could not connect to db")
}

// Ping db server to see if connection is alive and working
func (db *DB) Ping() error {
	return db.conn.Ping()
}
