package db

import (
	"log"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

type ScanResult struct {
	ScanIP    string
	Host      string
	CreatedAt time.Time
	RawData   []byte
}

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

func CreateSchema(databaseURL string) error {
	db, closeFn, err := Connect(databaseURL)
	if err != nil {
		return err
	}
	defer closeFn()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS scan_results(
		ID INT PRIMARY KEY     NOT NULL
	)`)
	return err
}
