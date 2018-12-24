package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSh(t *testing.T) {
	data, err := Sh("bash", "-c", "echo hello world!")
	assert.Nil(t, err)
	assert.Equal(t, "hello world!\n", string(data))
}
