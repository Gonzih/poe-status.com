package main

import (
	"fmt"
	"log"

	ping "github.com/sparrc/go-ping"
	yaml "gopkg.in/yaml.v2"
)

const portLimit = 996

type Host struct {
	Name      string
	HostName  string
	PingStats *ping.Statistics
	PingError error
	ScanData  []PortInfo
	ScanError error
}

func (h *Host) Debug() {
	data, err := yaml.Marshal(h)
	must(err)
	fmt.Println(string(data))
}

func (h *Host) Scan() {
	if h.ScanData != nil {
		return
	}

	h.ScanData, h.ScanError = Scan(h.HostName)
}

func (h *Host) IsUp() (bool, error) {
	h.Scan()
	if h.ScanError != nil {
		return false, h.ScanError
	}

	if len(h.ScanData) == 0 {
		return false, fmt.Errorf("Host %s is down", h.HostName)
	}

	openPorts := 0

	for _, port := range h.ScanData {
		if port.Open {
			openPorts++
			continue
		}

		// log.Printf("%d => %v", port.Port, port.Open)
	}

	if openPorts < portLimit {
		return false, fmt.Errorf("Host only had %d open ports", openPorts)
	}

	log.Printf("Host had %d open ports", openPorts)

	return true, nil
}
