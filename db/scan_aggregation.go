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

func intervalToSeconds(interval time.Duration) int {
	return int(interval.Seconds())
}

// AllScanAggregationsFor returns scan agggerations for last interval
func AllScanAggregationsFor(dbh dbHandler, interval time.Duration) ([]ScanAggr, error) {
	results := []ScanAggr{}
	err := dbh.Select(
		&results,
		`SELECT host,count(up) AS number_of_samples,up FROM scan_results WHERE created_at > NOW() - INTERVAL '1 second' * $1 GROUP BY host,up`,
		int(interval.Seconds()),
	)

	return results, err
}
