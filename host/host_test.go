package host

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/Gonzih/poe-status.com/rpc"
)

func TestScanError(t *testing.T) {
	req := &rpc.ScanResults{
		ScanError: "Nmap error something",
	}

	up, err := IsUp(req)
	assert.False(t, up)
	assert.NotNil(t, err)
}

func TestScanNoPorts(t *testing.T) {
	req := &rpc.ScanResults{}

	up, err := IsUp(req)
	assert.False(t, up)
	assert.NotNil(t, err)
}

func TestScanClosedPorts(t *testing.T) {
	ports := make([]*rpc.PortInfo, 160)
	for i := range ports {
		ports[i] = &rpc.PortInfo{
			Port: 1,
			Open: false,
		}
	}

	req := &rpc.ScanResults{
		Ports: ports,
	}

	up, err := IsUp(req)
	assert.False(t, up)
	assert.NotNil(t, err)
}

func TestScanOpenPorts(t *testing.T) {
	ports := make([]*rpc.PortInfo, 10)
	for i := range ports {
		ports[i] = &rpc.PortInfo{
			Port: 1,
			Open: true,
		}
	}

	req := &rpc.ScanResults{
		Ports: ports,
	}

	up, err := IsUp(req)
	assert.True(t, up)
	assert.Nil(t, err)
}
