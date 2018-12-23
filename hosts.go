package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

const portFailureThreshold = 150

type Host struct {
	Name      string
	HostName  string
	IsHostUp  bool
	ScanData  []PortInfo
	ScanError string
	UpdatedAt time.Time
}

func (h *Host) Debug() {
	data, err := yaml.Marshal(h)
	must(err)
	fmt.Println(string(data))
}

func (h *Host) ReScan(ports []int) {
	if time.Since(h.UpdatedAt) > time.Minute*5 {
		h.Scan(ports)
		up, err := h.IsUp()
		h.IsHostUp = up
		if h.ScanError == "" {
			h.ScanError = err.Error()
		}
		log.Printf("%s is %v: %v", h.Name, up, err)
	}
}

func (h *Host) Scan(ports []int) {
	if len(h.ScanData) > 0 {
		return
	}

	data, err := Scan(h.HostName, ports)
	h.ScanData = data
	h.ScanError = err.Error()
	h.UpdatedAt = time.Now()
}

func (h *Host) IsUp() (bool, error) {
	if h.ScanError != "" {
		return false, errors.New(h.ScanError)
	}

	if len(h.ScanData) == 0 {
		return false, fmt.Errorf("Host %s is down", h.HostName)
	}

	closedPorts := 0
	openPorts := 0

	for _, port := range h.ScanData {
		if !port.Open {
			log.Printf("Port %d is closed", port.Port)
			closedPorts++
		} else {
			openPorts++
		}
	}

	log.Printf("Host %s had %d closed and %d open ports", h.HostName, closedPorts, openPorts)

	if closedPorts > portFailureThreshold {
		return false, fmt.Errorf("Host %s had %d closed and %d open ports", h.HostName, closedPorts, openPorts)
	}

	return true, nil
}
