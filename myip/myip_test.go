package myip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyIp(t *testing.T) {
	ip, err := MyIP()
	assert.Nil(t, err)
	assert.NotNil(t, ip)
}
