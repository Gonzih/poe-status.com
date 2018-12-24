package scanner

import (
	nmap "github.com/tomsteele/go-nmap"
	"gitlab.com/Gonzih/poe-status.com/sh"
)

// NmapScan will scan given host using nmap shelling out
func NmapScan(host string) ([]PortInfo, error) {
	rawData, err := sh.Sh("nmap", "-oX", "-", host)
	if err != nil {
		return nil, err
	}

	nmapData, err := nmap.Parse(rawData)
	if err != nil {
		return nil, err
	}

	ports := make([]PortInfo, 0)

	for _, host := range nmapData.Hosts {
		for _, port := range host.Ports {
			p := PortInfo{Port: port.PortId, Open: port.State.State == "open"}
			ports = append(ports, p)
		}
	}

	return ports, nil
}
