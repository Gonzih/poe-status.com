package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const defaultUlimit = 1024

func Ulimit() int64 {
	out, err := sh("bash", "-c", "ulimit", "-n")

	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))
	if s == "unlimited" {
		log.Printf("Ulimit is unlimited, limiting to %d", defaultUlimit)
		return int64(defaultUlimit)
	}

	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	log.Printf("Ulimit is %d", i)

	return i
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func slurp(fname string) ([]byte, error) {
	return ioutil.ReadFile(fname)
}

func sh(args ...string) ([]byte, error) {
	log.Printf("$ %s", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}
