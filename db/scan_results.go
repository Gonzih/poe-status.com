package db

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type dbHandler interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	Queryx(query string, arg ...interface{}) (*sqlx.Rows, error)
}

// ScanResult represent record in a table
type ScanResult struct {
	ResultID  int       `db:"scan_result_id"`
	ScanIP    string    `db:"scan_ip"`
	Host      string    `db:"host"`
	Up        bool      `db:"up"`
	CreatedAt time.Time `db:"created_at"`
	QueryData []byte    `db:"query_data"`
}

// SaveScanResult saves one given scan result
func SaveScanResult(dbh dbHandler, sr *ScanResult) error {
	_, err := dbh.NamedExec(
		"INSERT INTO scan_results (scan_ip, host, up, created_at, query_data) VALUES (:scan_ip,:host,:up,:created_at,:query_data)",
		sr,
	)

	return err
}

// SelectScanResults selects scan results based on query
func SelectScanResults(dbh dbHandler, query string) ([]*ScanResult, error) {
	var results []*ScanResult
	res, err := dbh.Queryx(query)

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
func AllScanResults(dbh dbHandler) ([]*ScanResult, error) {
	return SelectScanResults(dbh, "SELECT * FROM scan_results")
}
