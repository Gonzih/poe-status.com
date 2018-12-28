package db

import (
	"fmt"
	"time"
)

// ScanResult represent record in a table
type ScanResult struct {
	ResultID  int       `db:"scan_result_id"`
	ScanIP    string    `db:"scan_ip"`
	Host      string    `db:"host"`
	Plaftorm  string    `db:"platform"`
	Up        bool      `db:"up"`
	CreatedAt time.Time `db:"created_at"`
	QueryData []byte    `db:"query_data"`
}

// ScanResultStore is a struct ot access scan results from db
type ScanResultStore struct {
	db dbHandler
}

// NewScanResultStore creates new store using given db handler
func NewScanResultStore(dbh dbHandler) *ScanResultStore {
	return &ScanResultStore{db: dbh}
}

// NewDefaultScanResultStore creates new store using default db connection
func NewDefaultScanResultStore() *ScanResultStore {
	return &ScanResultStore{db: DB()}
}

// SaveScanResult saves one given scan result
func (store *ScanResultStore) SaveScanResult(sr *ScanResult) error {
	_, err := store.db.NamedExec(
		"INSERT INTO scan_results (scan_ip,host,up,created_at,query_data,platform) VALUES (:scan_ip,:host,:up,:created_at,:query_data,:platform)",
		sr,
	)

	return err
}

// SelectScanResults selects scan results based on query
func (store *ScanResultStore) SelectScanResults(query string) ([]*ScanResult, error) {
	var results []*ScanResult
	res, err := store.db.Queryx(query)

	if err != nil {
		return results, err
	}

	for res.Next() {
		sr := &ScanResult{}
		err = res.StructScan(sr)
		if err != nil {
			return results, err
		}

		results = append(results, sr)
	}

	return results, nil
}

// AllScanResults returns all scan_results from db
func (store *ScanResultStore) AllScanResults() ([]*ScanResult, error) {
	return store.SelectScanResults("SELECT * FROM scan_results")
}

// DeleteResultsOlderThan will delete all results older than given age
func (store *ScanResultStore) DeleteResultsOlderThan(age time.Duration) error {
	if age < time.Hour*3 {
		return fmt.Errorf("Age should be older than 3 hours, got %v", age)
	}

	_, err := store.db.Exec(
		"DELETE FROM scan_results WHERE created_at < NOW() - INTERVAL '1 second' * $1",
		int(age.Seconds()),
	)

	return err
}
