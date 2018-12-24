package util

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gitlab.com/Gonzih/poe-status.com/sh"
)

const defaultUlimit = int64(1024)

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

	i, err := strconv.ParseInt(s, 10, 64)

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
