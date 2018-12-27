package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInto(t *testing.T) {
	var target int32
	in := []byte("10")
	err := parseInto(in, &target)
	assert.Nil(t, err)
	assert.Equal(t, target, int32(10))
}

func TestParseOutput(t *testing.T) {
	input := `PING google.com (172.217.1.14) 56(84) bytes of data.
64 bytes from iad23s25-in-f14.1e100.net (172.217.1.14): icmp_seq=1 ttl=53 time=4.07 ms
64 bytes from iad23s25-in-f14.1e100.net (172.217.1.14): icmp_seq=2 ttl=53 time=3.70 ms
64 bytes from iad23s25-in-f14.1e100.net (172.217.1.14): icmp_seq=3 ttl=53 time=4.31 ms
64 bytes from iad23s25-in-f14.1e100.net (172.217.1.14): icmp_seq=4 ttl=53 time=3.31 ms
64 bytes from iad23s25-in-f14.1e100.net (172.217.1.14): icmp_seq=5 ttl=53 time=3.69 ms

--- google.com ping statistics ---
100 packets transmitted, 90 received, 10% packet loss, time 8ms
rtt min/avg/max/mdev = 3.308/3.815/4.310/0.345 ms`

	out, err := parsePingOutput([]byte(input))
	assert.Nil(t, err)
	assert.Equal(t, out.transmitted, int32(100))
	assert.Equal(t, out.received, int32(90))
	assert.Equal(t, out.loss, int32(10))

	input = "PING google.com (172.217.0.238) 56(84) bytes of data.\n64 bytes from dfw06s38-in-f14.1e100.net (172.217.0.238): icmp_seq=1 ttl=53 time=1.90 ms\n64 bytes from dfw06s38-in-f14.1e100.net (172.217.0.238): icmp_seq=2 ttl=53 time=3.80 ms\n64 bytes from dfw06s38-in-f14.1e100.net (172.217.0.238): icmp_seq=3 ttl=53 time=5.71 ms\n64 bytes from dfw06s38-in-f14.1e100.net (172.217.0.238): icmp_seq=4 ttl=53 time=2.100 ms\n64 bytes from dfw06s38-in-f14.1e100.net (172.217.0.238): icmp_seq=5 ttl=53 time=4.99 ms\n64 bytes from dfw06s38-in-f14.1e100.net (172.217.0.238): icmp_seq=6 ttl=53 time=3.55 ms\n\n--- google.com ping statistics ---\n6 packets transmitted, 6 received, 0% packet loss, time 10ms\nrtt min/avg/max/mdev = 1.897/3.822/5.712/1.253 ms\n"

	out, err = parsePingOutput([]byte(input))
	assert.Nil(t, err)
	assert.Equal(t, out.transmitted, int32(6))
	assert.Equal(t, out.received, int32(6))
	assert.Equal(t, out.loss, int32(0))
}

func TestPingSimple(t *testing.T) {
	out, err := Ping("google.com", 6)
	assert.Nil(t, err)
	assert.NotEqual(t, out.Output, "")
	assert.Equal(t, out.Transmitted, int32(6))
	assert.Equal(t, out.Received, int32(6))
	assert.Equal(t, out.Loss, int32(0))
}
