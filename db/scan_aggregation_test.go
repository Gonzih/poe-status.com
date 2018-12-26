package db

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestIntervalToSeconds(t *testing.T) {
	samples := map[time.Duration]int{
		time.Hour:        3600,
		time.Minute * 15: 900,
	}

	for input, output := range samples {
		s := intervalToSeconds(input)
		assert.Equal(t, s, output)
	}
}

func TestScanAggregationsInTimerange(t *testing.T) {
	dbUp(t)

	withTestTransaction(func(tx *sqlx.Tx) {
		result1 := &ScanResult{
			ScanIP:    "192.168.2.1",
			Host:      "login.pathoexile.com",
			Up:        true,
			CreatedAt: time.Now().Add(time.Minute * -10),
			QueryData: []byte("{}"),
		}

		err := SaveScanResult(tx, result1)
		assert.Nil(t, err)

		results, err := AllScanAggregationsFor(tx, time.Minute*15)
		assert.Nil(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, results[0].Host, result1.Host)
		assert.Equal(t, results[0].Up, true)
		assert.Equal(t, results[0].NumberOfSamples, 1)
	})
}

func TestScanAggregationsOutOfTimerange(t *testing.T) {
	dbUp(t)

	withTestTransaction(func(tx *sqlx.Tx) {
		result1 := &ScanResult{
			ScanIP:    "192.168.2.1",
			Host:      "login.pathoexile.com",
			Up:        true,
			CreatedAt: time.Now().Add(time.Hour * -10),
			QueryData: []byte("{}"),
		}

		err := SaveScanResult(tx, result1)
		assert.Nil(t, err)

		results, err := AllScanAggregationsFor(tx, time.Minute*15)
		assert.Nil(t, err)
		assert.Len(t, results, 0)
	})
}
