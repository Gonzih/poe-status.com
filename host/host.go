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

func isUpBasedOnPing(scan *rpc.ScanResults) (bool, error) {
	if scan.PingInfo == nil {
		return false, fmt.Errorf("No ping info found for %s", scan.Host)
	}

	if scan.PingInfo.Received < scan.PingInfo.Transmitted || scan.PingInfo.Loss > int32(10) {
		return false, fmt.Errorf(
			"Host %s lost some packets: %d received, %d transmitted, %d loss",
			scan.Host,
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
		pcUp, pcErr := isPCUp(scan)
		pingUp, pingErr := isUpBasedOnPing(scan)
		var err error
		if pcErr != nil {
			err = pcErr
		} else if pingErr != nil {
			err = pingErr
		}

		return pcUp && pingUp, err
	}

	if scan.Platform == "XBOX" {
		return isUpBasedOnPing(scan)
	}

	return false, errors.New("Nothing happened")
}
