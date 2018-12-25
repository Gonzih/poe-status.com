package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/Gonzih/poe-status.com/migrations"
)

func dbUp(t *testing.T) *sql.Tx {
	err := migrations.Up("../migrations", TestDBURL())
	assert.Nil(t, err)
	Init(TestDBURL())
	tx, err := db.Begin()
	assert.Nil(t, err)

	return tx
}

func dbDown() {
	err := migrations.Down("../migrations", TestDBURL())
	if err != nil {
		panic(err)
	}
}

func TestDBConnection(t *testing.T) {
	dbUp(t)
	defer dbDown()

	db, closeFn, err := Connect(TestDBURL())
	assert.Nil(t, err)
	assert.NotNil(t, db)
	defer closeFn()

	result, err := db.Query("SELECT * FROM scan_results")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestDBCreateSchema(t *testing.T) {
	dbUp(t)
	dbDown()
}
