package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadYaml(t *testing.T) {
	cfg, err := ReadYAML()
	assert.Nil(t, err)
	assert.Len(t, cfg.PC, 13)
	assert.Len(t, cfg.XBOX, 8)
	assert.Len(t, cfg.Ports, 996)
}

func TestAllHosts(t *testing.T) {
	cfg, err := ReadYAML()
	assert.Nil(t, err)
	assert.Len(t, cfg.AllHosts(), 21)
}

func TestAllHostsOrder(t *testing.T) {
	cfg, err := ReadYAML()
	assert.Nil(t, err)
	assert.Equal(t, cfg.AllHosts()[0].Name, "Amsterdam")
}

func TestMainToken(t *testing.T) {
	assert.Equal(t, MainToken(), "TESTTOKEN")
}
