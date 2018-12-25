package host

import (
	"errors"
	"fmt"
	"log"

	"gitlab.com/Gonzih/poe-status.com/rpc"
)

const portFailureThreshold = 150

func IsUp(scan *rpc.ScanResults) (bool, error) {
	if scan.ScanError != "" {
		return false, errors.New(scan.ScanError)
	}

	if len(scan.Ports) == 0 {
		return false, fmt.Errorf("Host %s is down since no ports were open", scan.Host)
	}

	closedPorts := 0
	openPorts := 0

	for _, port := range scan.Ports {
		if !port.Open {
			// log.Printf("Port %d is closed", port.Port)
			closedPorts++
		} else {
			openPorts++
		}
	}

	msg := fmt.Sprintf("Host %s had %d closed and %d open ports", scan.Host, closedPorts, openPorts)

	if closedPorts > portFailureThreshold {
		return false, errors.New(msg)
	}

	log.Println(msg)
	return true, nil
}
