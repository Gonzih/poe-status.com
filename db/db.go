package db

import (
	"log"
	"os"

	"database/sql"

	// imported as psql driver for database/sql
	_ "github.com/lib/pq"
)

var db *sql.DB

// DB exposes db instance to the outer world
func DB() *sql.DB {
	if db == nil {
		log.Fatalf("DB was not initialized")
	}

	return db
}

// Init will initalize global connection
func Init(databaseURL string) error {
	if db == nil {
		tmpdb, _, err := Connect(databaseURL)
		if err != nil {
			return err
		}
		db = tmpdb
	}

	return nil
}

// Connect connects to the database
func Connect(databaseURL string) (*sql.DB, func(), error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, func() {}, err
	}

	return db, func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing db: %s", err)
		}
	}, nil
}

// TestDBURL returns database URL string based on the environment variables
func TestDBURL() string {
	e := os.Getenv("TEST_DATABASE_URL")
	if e != "" {
		return e

	}

	return "postgres://postgres@localhost/poetest?sslmode=disable"
}
