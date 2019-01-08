package scanner

import (
	"fmt"
	"regexp"

	"github.com/Gonzih/poe-status.com/rpc"
	"github.com/Gonzih/poe-status.com/sh"
	"github.com/Gonzih/poe-status.com/util"
)

type pingOutput struct {
	transmitted int32
	received    int32
	loss        int32
}

func parseInto(in []byte, target *int32) error {
	s := string(in)
	parsed, err := util.ParseInt32(s)
	if err != nil {
		return fmt.Errorf("Could not parse %s in to int32: %s", in, err)
	}
	*target = parsed

	return nil
}

var pingOutputRegexp = regexp.MustCompile(`(\d+) packets transmitted, (\d+) received, (\d+)% packet loss,`)

func parsePingOutput(input []byte) (pingOutput, error) {
	out := pingOutput{}

	matches := pingOutputRegexp.FindSubmatch(input)

	if len(matches) < 4 {
		return out, fmt.Errorf(`Could not find ping data (len is %d) in input "%s"`, len(matches), string(input))
	}

	err := parseInto(matches[1], &out.transmitted)
	if err != nil {
		return out, err
	}

	err = parseInto(matches[2], &out.received)
	if err != nil {
		return out, err
	}

	err = parseInto(matches[3], &out.loss)
	if err != nil {
		return out, err
	}

	return out, nil
}

// Ping will run ping against the host
func Ping(host string, n int) (*rpc.PingInfo, error) {
	out, err := sh.Sh("ping", "-c", fmt.Sprintf("%d", n), host)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, string(out))
	}

	res, err := parsePingOutput(out)
	if err != nil {
		return nil, err
	}

	return &rpc.PingInfo{
		Output:      string(out),
		Transmitted: res.transmitted,
		Received:    res.received,
		Loss:        res.loss,
	}, nil
}
