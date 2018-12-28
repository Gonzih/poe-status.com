package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	// imported as psql driver for database/sql
	_ "github.com/lib/pq"
)

type dbHandler interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, arg ...interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	Queryx(query string, arg ...interface{}) (*sqlx.Rows, error)
	Select(dest interface{}, query string, args ...interface{}) error
}

var db *sqlx.DB

// DB exposes db instance to the outer world
func DB() *sqlx.DB {
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
func Connect(databaseURL string) (*sqlx.DB, func(), error) {
	db, err := sqlx.Open("postgres", databaseURL)
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
