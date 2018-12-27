package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gitlab.com/Gonzih/poe-status.com/app/config"
	"gitlab.com/Gonzih/poe-status.com/db"
	"gitlab.com/Gonzih/poe-status.com/util"
)

type aggregationWrapper struct {
	ServerName   string       `json:"server_name"`
	Platform     string       `json:"platform"`
	UpEvidence   *db.ScanAggr `json:"up_evidence,omitempty"`
	DownEvidence *db.ScanAggr `json:"down_evidence,omitempty"`
}

var _cfg *config.Config

func cfg() *config.Config {
	if _cfg != nil {
		return _cfg
	}

	c, err := config.ReadYAML()
	util.Must(err)
	_cfg = c

	return _cfg
}

// ScanAggrEndpoint is json endpoint that returns aggregated scan data
func ScanAggrEndpoint(w http.ResponseWriter, r *http.Request) {
	aggregations, err := db.AllScanAggregationsFor(db.DB(), time.Hour*3)
	if err != nil {
		log.Fatal(err)
	}
	allHosts := cfg().AllHosts()

	response := make([]aggregationWrapper, len(allHosts))

	for i, host := range allHosts {
		for _, aggr := range aggregations {
			if host.Host == aggr.Host {
				response[i].ServerName = host.Name
				response[i].Platform = host.Platform
				if aggr.Up {
					response[i].UpEvidence = &aggr
				} else {
					response[i].DownEvidence = &aggr
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
