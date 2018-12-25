package db

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dbURL() string {
	e := os.Getenv("TEST_DATABASE_URL")
	if e != "" {
		return e

	}

	return "postgres://postgres@localhost/poetest?sslmode=disable"
}

func TestDBConnection(t *testing.T) {
	db, closeFn, err := Connect(dbURL())
	assert.Nil(t, err)
	assert.NotNil(t, db)
	defer closeFn()

	result, err := db.Query("SELECT true")
	log.Println(result)
}

func TestDBCreateSchema(t *testing.T) {
	err := CreateSchema(dbURL())
	assert.Nil(t, err)
}
