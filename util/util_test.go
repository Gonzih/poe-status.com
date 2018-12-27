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

func TestParseInt32(t *testing.T) {
	n, err := ParseInt32("123")
	assert.Nil(t, err)
	assert.Equal(t, n, int32(123))
}

func TestParseInt64(t *testing.T) {
	n, err := ParseInt64("123")
	assert.Nil(t, err)
	assert.Equal(t, n, int64(123))
}
