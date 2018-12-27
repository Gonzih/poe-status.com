package db

import (
	"time"
)

// ScanAggr represents struct for scan aggregations over some interval
type ScanAggr struct {
	Host            string `db:"host" json:"host"`
	NumberOfSamples int    `db:"number_of_samples" json:"number_of_samples"`
	Up              bool   `db:"up" json:"up"`
}

// ScanAggrStore is a struct ot access agggeration related db stuff
type ScanAggrStore struct {
	db dbHandler
}

// NewScanAggrStore creates new store using given db handler
func NewScanAggrStore(dbh dbHandler) *ScanAggrStore {
	return &ScanAggrStore{db: dbh}
}

// NewDefaultScanAggrStore creates new store using default db connection
func NewDefaultScanAggrStore() *ScanAggrStore {
	return &ScanAggrStore{db: DB()}
}

// AllScanAggregationsFor returns scan agggerations for last interval
func (store *ScanAggrStore) AllScanAggregationsFor(interval time.Duration) ([]ScanAggr, error) {
	results := []ScanAggr{}
	err := store.db.Select(
		&results,
		`SELECT host,count(up) AS number_of_samples,up FROM scan_results WHERE created_at > NOW() - INTERVAL '1 second' * $1 GROUP BY host,up`,
		int(interval.Seconds()),
	)

	return results, err
}

func intervalToSeconds(interval time.Duration) int {
	return int(interval.Seconds())
}
