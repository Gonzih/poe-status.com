package db

import "time"

// ScanResult represent record in a table
type ScanResult struct {
	ResultID  int
	ScanIP    string
	Host      string
	CreatedAt time.Time
	RawData   []byte
}

// SaveScanResult saves one given scan result
func SaveScanResult(sr *ScanResult) error {
	ins, err := db.Prepare(
		"INSERT INTO scan_results (scan_ip, host, created_at, query_data) VALUES ($1,$2,$3,$4)",
	)
	if err != nil {
		return err
	}

	_, err = ins.Exec(
		sr.ScanIP,
		sr.Host,
		sr.CreatedAt,
		sr.RawData,
	)

	return err
}

// SelectScanResults selects scan results based on query
func SelectScanResults(query string) ([]*ScanResult, error) {
	var results []*ScanResult
	res, err := db.Query(query)

	if err != nil {
		return results, err
	}

	for res.Next() {
		sr := &ScanResult{}
		err = res.Scan(
			&sr.ResultID,
			&sr.ScanIP,
			&sr.Host,
			&sr.CreatedAt,
			&sr.RawData,
		)
		if err != nil {
			return results, err
		}

		results = append(results, sr)
	}

	return results, nil
}

// AllScanResults returns all scan_results from db
func AllScanResults() ([]*ScanResult, error) {
	return SelectScanResults("SELECT * FROM scan_results")
}
