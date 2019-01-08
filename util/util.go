package util

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/Gonzih/poe-status.com/sh"
)

const defaultUlimit = int64(1024)

// ParseInt32 will try to parse integer in to int32 type
func ParseInt32(in string) (int32, error) {
	i, err := strconv.ParseInt(in, 10, 64)

	return int32(i), err
}

// ParseInt64 will try to parse integer in to int64 type
func ParseInt64(in string) (int64, error) {
	return strconv.ParseInt(in, 10, 64)
}

// Ulimit will try to get ulimit on a system or return a default var
func Ulimit() int64 {
	out, err := sh.Sh("bash", "-c", "ulimit", "-n")

	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))
	if s == "unlimited" {
		log.Printf("Ulimit is unlimited, limiting to %d", defaultUlimit)
		return defaultUlimit
	}

	i, err := ParseInt64(s)

	if err != nil {
		panic(err)
	}

	log.Printf("Ulimit is %d", i)

	return i
}

// Must will panic on error
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Slurp will read entire file in to memory
func Slurp(fname string) ([]byte, error) {
	return ioutil.ReadFile(fname)
}
