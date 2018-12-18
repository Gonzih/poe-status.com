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

func (h *Host) Scan(ports []int) {
	if h.ScanData != nil {
		return
	}

	h.ScanData, h.ScanError = Scan(h.HostName, ports)
}

func (h *Host) IsUp() (bool, error) {
	if h.ScanError != nil {
		return false, h.ScanError
	}

	if len(h.ScanData) == 0 {
		return false, fmt.Errorf("Host %s is down", h.HostName)
	}

	for _, port := range h.ScanData {
		if !port.Open {
			log.Printf("Port %d is down", port.Port)
			return false, fmt.Errorf("Port %d on host %s is closed", port, h.HostName)
		}
	}

	return true, nil
}
