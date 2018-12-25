package migrations

import (
	"errors"
	"testing"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/stretchr/testify/assert"
	"gitlab.com/Gonzih/poe-status.com/db"
)

func TestSQLMigrationsFolderPath(t *testing.T) {
	path := SQLMigrationsFolderPath(".")
	assert.Contains(t, path, "migrations/sql")
}

func TestFormatSqlErrorNil(t *testing.T) {
	err := FormatSQLError(nil)
	assert.Nil(t, err)
}

func TestFormatSqlErrorStringErr(t *testing.T) {
	err := FormatSQLError(errors.New("Test error"))
	assert.NotNil(t, err)
}

func TestFormatSqlErrorDBErr(t *testing.T) {
	err := FormatSQLError(database.Error{
		Line:  1,
		Query: []byte("query sample"),
		Err:   "Wrong syntax",
	})
	assert.NotNil(t, err)
}

func TestDBMigrationsUpAndDown(t *testing.T) {
	url := db.TestDBURL()
	err := Up(".", url)
	assert.Nil(t, err)
	err = Down(".", url)
	assert.Nil(t, err)
}
