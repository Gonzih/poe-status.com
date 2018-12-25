package db

import (
	"log"
	"os"
	"time"

	"database/sql"

	// imported as psql driver for database/sql
	_ "github.com/lib/pq"
)

// ScanResult represent record in a table
type ScanResult struct {
	ScanIP    string
	Host      string
	CreatedAt time.Time
	RawData   []byte
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
