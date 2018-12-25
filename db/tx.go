package db

import (
	"database/sql"
	"errors"
	"log"
)

var RollbackError = errors.New("rollback this tx")

func newTx() *sql.Tx {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error starting new transaction: %s", err)
	}

	return tx
}

// WithTransaction runs given function in a transaction
func WithTransaction(f func(*sql.Tx) error) error {
	tx := newTx()
	ferr := f(tx)

	if ferr != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}

		if ferr == RollbackError {
			return nil
		}

		return ferr
	}

	return tx.Commit()
}
