package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlurp(t *testing.T) {
	data, err := Slurp("fixtures/data.txt")
	assert.Nil(t, err)
	assert.Equal(t, "hello world!\n", string(data))
}

func TestMustDoesNothingOnNil(t *testing.T) {
	Must(nil)
}

func TestMustPanicsOnErr(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()

	Must(errors.New("test error"))
}

func TestUlimit(t *testing.T) {
	limit := Ulimit()
	assert.Equal(t, defaultUlimit, limit)
}
