package db

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScanResultSaveNoTx(t *testing.T) {
	dbUp(t)
	defer dbDown()

	result1 := &ScanResult{
		ScanIP:    "192.168.2.1",
		Host:      "login.pathoexile.com",
		CreatedAt: time.Now(),
		RawData:   []byte("{}"),
	}

	err := SaveScanResult(db, result1)
	assert.Nil(t, err)
}

func TestScanResultSaveInTx(t *testing.T) {
	dbUp(t)

	WithTransaction(func(tx *sql.Tx) error {
		result1 := &ScanResult{
			ScanIP:    "192.168.2.1",
			Host:      "login.pathoexile.com",
			Up:        true,
			CreatedAt: time.Now(),
			RawData:   []byte("{}"),
		}

		err := SaveScanResult(tx, result1)
		assert.Nil(t, err)

		return RollbackError
	})
}

func TestScanResultSaveLoad(t *testing.T) {
	dbUp(t)

	WithTransaction(func(tx *sql.Tx) error {
		result1 := &ScanResult{
			ScanIP:    "192.168.2.1",
			Host:      "login.pathoexile.com",
			CreatedAt: time.Now(),
			RawData:   []byte("{}"),
		}

		err := SaveScanResult(tx, result1)
		assert.Nil(t, err)

		results, err := AllScanResults(tx)
		assert.Nil(t, err)
		assert.Len(t, results, 1)

		result2 := results[0]

		assert.Equal(t, result1.ScanIP, result2.ScanIP)
		assert.Equal(t, result1.Host, result2.Host)
		assert.True(t, result1.CreatedAt.Sub(result2.CreatedAt) < time.Second)

		return RollbackError
	})
}
