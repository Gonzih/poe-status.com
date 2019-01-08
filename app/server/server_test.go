package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Gonzih/poe-status.com/rpc"
)

func dummyServer() *PoeStatusServer {
	return &PoeStatusServer{}
}

func TestNoTokenPresent(t *testing.T) {
	_, err := dummyServer().SaveScanResults(context.TODO(), &rpc.ScanResults{})
	assert.NotNil(t, err)
}

func TestInvalidToken(t *testing.T) {
	_, err := dummyServer().SaveScanResults(context.TODO(), &rpc.ScanResults{
		AuthToken: &rpc.AuthToken{
			Token: "bebebe",
		},
	})
	assert.NotNil(t, err)
}
