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

// ScanAggrEndpoint is json endpoint that returns aggregated scan data
func ScanAggrEndpoint(w http.ResponseWriter, r *http.Request) {
	aggregations, err := db.AllScanAggregationsFor(db.DB(), time.Hour*3)
	if err != nil {
		log.Fatal(err)
	}
	allHosts := cfg().AllHosts()

	response := make([]aggregationWrapper, 0)
	tmp := make(map[string]aggregationWrapper, 0)

	for _, host := range allHosts {
		wrapper, ok := tmp[host.Host]
		if !ok {
			wrapper = aggregationWrapper{}
		}
		wrapper.ServerName = host.Name
		wrapper.Host = host.Host
		wrapper.Platform = host.Platform

		for _, aggr := range aggregations {
			if wrapper.Host == aggr.Host {
				if aggr.Up {
					wrapper.UpEvidence = aggr
				} else {
					wrapper.DownEvidence = aggr
				}
			}
		}

		tmp[host.Host] = wrapper
	}

	log.Println(tmp)

	for _, host := range allHosts {
		response = append(response, tmp[host.Host])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
