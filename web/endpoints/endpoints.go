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
	ServerName string      `json:"server_name"`
	Data       db.ScanAggr `json:"data"`
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

func ScanAggrEndpoint(w http.ResponseWriter, r *http.Request) {
	aggregations, err := db.AllScanAggregationsFor(db.DB(), time.Hour*15)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]aggregationWrapper, len(aggregations))

	for i, aggr := range aggregations {
		for host, name := range cfg().PC {
			if host == aggr.Host {
				data[i].ServerName = name
				data[i].Data = aggr
				break
			}
		}
	}

	json.NewEncoder(w).Encode(data)
}
