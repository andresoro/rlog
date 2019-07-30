package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	migrate "github.com/rubenv/sql-migrate"

	_ "github.com/lib/pq"
)

// Config holds postgres related variables
type Config struct {
	User      string
	Pass      string
	DB        string
	Migration string
}

// DB wrap around an sql connection
type DB struct {
	conn   *sql.DB
	Config *Config
}

// Connect to postgres db, uses ENV variables from docker image
func (db *DB) Connect() error {

	s := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, db.Config.User, db.Config.Pass, db.Config.DB)
	conn, err := sql.Open("postgres", s)
	if err != nil {
		return err
	}

	db.conn = conn

	log.Println("Attempting to register connection with postgres db...")
	log.Println("Will abort if connection fails after 30 seconds")
	for i := 0; i < 15; i++ {
		err := db.Ping()
		if err == nil { // if connection established
			return nil
		}
		time.Sleep(2 * time.Second)
	}

	return errors.New("Could not connect to db")
}

// Migrate schema changes to postgres on connect
func (db *DB) Migrate() error {
	migrations := &migrate.FileMigrationSource{
		Dir: db.Config.Migration,
	}

	n, err := migrate.Exec(db.conn, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	if n > 0 {
		log.Printf("Successfully applied %d migrations\n", n)
	} else {
		log.Println("No schema changes - no migrations needed")
	}

	return nil
}

// Ping db server to see if connection is alive and working
func (db *DB) Ping() error {
	return db.conn.Ping()
}

// Close the connection to postgres
func (db *DB) Close() error {
	return db.conn.Close()
}
