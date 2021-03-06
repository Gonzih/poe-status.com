package host

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Gonzih/poe-status.com/rpc"
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
		Ports:    ports,
		Platform: "PC",
		PingInfo: &rpc.PingInfo{
			Transmitted: 100,
			Received:    100,
			Loss:        0,
		},
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
		Ports:    ports,
		Platform: "PC",
		PingInfo: &rpc.PingInfo{
			Transmitted: 100,
			Received:    100,
			Loss:        0,
		},
	}

	up, err := IsUp(req)
	assert.True(t, up)
	assert.Nil(t, err)
}

func TestScanOpenPortsAndFailedPing(t *testing.T) {
	ports := make([]*rpc.PortInfo, 10)
	for i := range ports {
		ports[i] = &rpc.PortInfo{
			Port: 1,
			Open: true,
		}
	}

	req := &rpc.ScanResults{
		Ports:    ports,
		Platform: "PC",
		PingInfo: &rpc.PingInfo{
			Transmitted: 100,
			Received:    90,
			Loss:        10,
		},
	}

	up, err := IsUp(req)
	assert.False(t, up)
	assert.NotNil(t, err)
}

func TestPingSuccess(t *testing.T) {
	pingInfo := &rpc.PingInfo{
		Transmitted: 100,
		Received:    100,
		Loss:        0,
	}

	req := &rpc.ScanResults{
		PingInfo: pingInfo,
		Platform: "XBOX",
	}

	up, err := IsUp(req)
	assert.True(t, up)
	assert.Nil(t, err)
}

func TestPingSuccessWithSomeLoss(t *testing.T) {
	pingInfo := &rpc.PingInfo{
		Transmitted: 100,
		Received:    100,
		Loss:        5,
	}

	req := &rpc.ScanResults{
		PingInfo: pingInfo,
		Platform: "XBOX",
	}

	up, err := IsUp(req)
	assert.True(t, up)
	assert.Nil(t, err)
}

func TestPingFaildueToLoss(t *testing.T) {
	pingInfo := &rpc.PingInfo{
		Transmitted: 100,
		Received:    90,
		Loss:        10,
	}

	req := &rpc.ScanResults{
		PingInfo: pingInfo,
		Platform: "XBOX",
	}

	up, err := IsUp(req)
	assert.False(t, up)
	assert.NotNil(t, err)
}
