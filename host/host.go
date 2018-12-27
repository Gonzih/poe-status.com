package host

import (
	"errors"
	"fmt"
	"log"

	"gitlab.com/Gonzih/poe-status.com/rpc"
)

const portFailureThreshold = 150

func isPCUp(scan *rpc.ScanResults) (bool, error) {
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

func isXboxUp(scan *rpc.ScanResults) (bool, error) {
	if scan.PingInfo == nil {
		return false, errors.New("No ping info found")
	}

	if scan.PingInfo.Received < scan.PingInfo.Transmitted || scan.PingInfo.Loss > int32(0) {
		return false, fmt.Errorf(
			"Lost some packages: %d transmitted, %d received, %d loss",
			scan.PingInfo.Received,
			scan.PingInfo.Transmitted,
			scan.PingInfo.Loss,
		)
	}

	log.Printf("Ping results look fine for %s", scan.Host)
	return true, nil
}

// IsUp determines if host is up
func IsUp(scan *rpc.ScanResults) (bool, error) {
	if scan.ScanError != "" {
		return false, errors.New(scan.ScanError)
	}

	if scan.Platform == "PC" {
		return isPCUp(scan)
	}

	if scan.Platform == "XBOX" {
		return isXboxUp(scan)
	}

	return false, errors.New("Nothing happened")
}
