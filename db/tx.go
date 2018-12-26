package db

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

var RollbackError = errors.New("rollback this tx")

func newTx() *sqlx.Tx {
	tx, err := db.Beginx()
	if err != nil {
		log.Fatalf("Error starting new transaction: %s", err)
	}

	return tx
}

// WithTransaction runs given function in a transaction
func WithTransaction(f func(*sqlx.Tx) error) error {
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
