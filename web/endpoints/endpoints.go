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
	ServerName   string      `json:"server_name"`
	Platform     string      `json:"platform"`
	Host         string      `json:"host"`
	UpEvidence   db.ScanAggr `json:"up_evidence,omitempty"`
	DownEvidence db.ScanAggr `json:"down_evidence,omitempty"`
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

func getAggrData() []aggregationWrapper {
	aggStore := db.NewDefaultScanAggrStore()
	aggregations, err := aggStore.AllScanAggregationsFor(time.Minute * 15)
	if err != nil {
		log.Fatal(err)
	}
	allHosts := cfg().AllHosts()

	response := make([]aggregationWrapper, len(allHosts))

	for i, host := range allHosts {
		response[i].ServerName = host.Name
		response[i].Host = host.Host
		response[i].Platform = host.Platform

		for _, aggr := range aggregations {
			if response[i].Host == aggr.Host {
				if aggr.Up {
					response[i].UpEvidence = aggr
				} else {
					response[i].DownEvidence = aggr
				}
			}
		}
	}

	return response
}

// ScanAggrEndpoint is json endpoint that returns aggregated scan data
func ScanAggrEndpoint(w http.ResponseWriter, r *http.Request) {
	response := getAggrData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
