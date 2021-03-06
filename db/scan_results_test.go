package db

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func withTestTransaction(f func(tx *sqlx.Tx)) {
	WithTransaction(func(tx *sqlx.Tx) error {
		f(tx)
		return RollbackError
	})
}

func TestScanResultSaveNoTx(t *testing.T) {
	dbUp(t)
	defer dbDown()

	result1 := &ScanResult{
		ScanIP:    "192.168.2.1",
		Host:      "login.pathoexile.com",
		CreatedAt: time.Now(),
		QueryData: []byte("{}"),
		Plaftorm:  "PC",
	}

	resultStore := NewDefaultScanResultStore()
	err := resultStore.SaveScanResult(result1)
	assert.Nil(t, err)
}

func TestScanResultSaveInTx(t *testing.T) {
	dbUp(t)

	withTestTransaction(func(tx *sqlx.Tx) {
		result1 := &ScanResult{
			ScanIP:    "192.168.2.1",
			Host:      "login.pathoexile.com",
			Up:        true,
			CreatedAt: time.Now(),
			QueryData: []byte("{}"),
			Plaftorm:  "PC",
		}

		resultStore := NewScanResultStore(tx)
		err := resultStore.SaveScanResult(result1)
		assert.Nil(t, err)
	})
}

func TestScanResultSaveLoad(t *testing.T) {
	dbUp(t)

	withTestTransaction(func(tx *sqlx.Tx) {
		result1 := &ScanResult{
			ScanIP:    "192.168.2.1",
			Host:      "login.pathoexile.com",
			CreatedAt: time.Now(),
			QueryData: []byte("{}"),
			Plaftorm:  "PC",
		}

		resultStore := NewScanResultStore(tx)
		err := resultStore.SaveScanResult(result1)
		assert.Nil(t, err)

		results, err := resultStore.AllScanResults()
		assert.Nil(t, err)
		assert.Len(t, results, 1)

		result2 := results[0]

		assert.Equal(t, result1.ScanIP, result2.ScanIP)
		assert.Equal(t, result1.Host, result2.Host)
		assert.True(t, result1.CreatedAt.Sub(result2.CreatedAt) < time.Second)
	})
}

func TestDeleteResultsOlderThan(t *testing.T) {
	dbUp(t)

	withTestTransaction(func(tx *sqlx.Tx) {
		result1 := &ScanResult{
			ScanIP:    "192.168.2.1",
			Host:      "login.pathoexile.com",
			CreatedAt: time.Now(),
			QueryData: []byte("{}"),
			Plaftorm:  "PC",
		}

		result2 := &ScanResult{
			ScanIP:    "192.168.2.3",
			Host:      "au.xbox.pathoexile.com",
			CreatedAt: time.Now().Add(time.Hour * -12),
			QueryData: []byte("{}"),
			Plaftorm:  "Xbox",
		}

		resultStore := NewScanResultStore(tx)
		err := resultStore.SaveScanResult(result1)
		assert.Nil(t, err)
		err = resultStore.SaveScanResult(result2)
		assert.Nil(t, err)

		err = resultStore.DeleteResultsOlderThan(time.Hour * 9)
		assert.Nil(t, err)

		results, err := resultStore.AllScanResults()
		assert.Nil(t, err)
		assert.Len(t, results, 1)

		result3 := results[0]

		assert.Equal(t, result1.ScanIP, result3.ScanIP)
		assert.Equal(t, result1.Host, result3.Host)
		assert.True(t, result1.CreatedAt.Sub(result3.CreatedAt) < time.Second)
	})
}

func TestDeleteResultsOlderThanShouldNotAllowSmallDuration(t *testing.T) {
	dbUp(t)
	err := NewDefaultScanResultStore().DeleteResultsOlderThan(time.Hour * 1)
	assert.NotNil(t, err)
}
