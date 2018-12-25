package db

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/Gonzih/poe-status.com/migrations"
)

func dbUp(t *testing.T) {
	err := migrations.Up("../migrations", TestDBURL())
	assert.Nil(t, err)
}

func dbDown(t *testing.T) {
	err := migrations.Up("../migrations", TestDBURL())
	assert.Nil(t, err)
}

func TestDBConnection(t *testing.T) {
	dbUp(t)
	defer dbDown(t)

	db, closeFn, err := Connect(TestDBURL())
	assert.Nil(t, err)
	assert.NotNil(t, db)
	defer closeFn()

	result, err := db.Query("SELECT true")
	log.Println(result)
}

func TestDBCreateSchema(t *testing.T) {
	dbUp(t)
	dbDown(t)
}
