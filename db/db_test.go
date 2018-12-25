package db

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	db, closeFn, err := Connect(TestDBURL())
	assert.Nil(t, err)
	assert.NotNil(t, db)
	defer closeFn()

	result, err := db.Query("SELECT true")
	log.Println(result)
}

func TestDBCreateSchema(t *testing.T) {
	err := CreateSchema(TestDBURL())
	assert.Nil(t, err)
}
