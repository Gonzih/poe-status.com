package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithTransactionDummyCommit(t *testing.T) {
	err := WithTransaction(func(tx *sql.Tx) error {
		return nil
	})

	assert.Nil(t, err)

}

func TestWithTransactionDummyRollback(t *testing.T) {
	err := WithTransaction(func(tx *sql.Tx) error {
		return RollbackError
	})

	assert.Nil(t, err)
}

func TestWithTransactionDummyError(t *testing.T) {
	err := WithTransaction(func(tx *sql.Tx) error {
		return errors.New("some scary exception")
	})

	assert.NotNil(t, err)
}
