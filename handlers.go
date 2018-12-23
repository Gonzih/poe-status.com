package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var hostState sync.Map

func StatusHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := readServersYaml()
	// ctx := r.Context()

	hosts := make([]*Host, 0)
	for name, host := range data.Pc {
		var h Host

		item, ok := hostState.Load(host)
		if ok {
			h, ok = item.(Host)
			if !ok {
				log.Fatalf("Unable to convert %v to type *Host", item)
			}
		} else {
			h = Host{Name: name, HostName: host}
		}

		h.ReScan(data.Ports)
		hostState.Store(host, h)
		hosts = append(hosts, h)
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(hosts)
	must(err)
}
