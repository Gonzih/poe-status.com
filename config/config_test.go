package config

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DataDir http.FileSystem = http.Dir("data")

func TestReadYaml(t *testing.T) {
	cfg, err := ReadYAML()
	assert.Nil(t, err)
	assert.Len(t, cfg.PC, 13)
	assert.Len(t, cfg.XBOX, 8)
	assert.Len(t, cfg.Ports, 996)
}
