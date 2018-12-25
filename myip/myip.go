package myip

import (
	"io/ioutil"
	"net"
	"net/http"
)

var queryURL = "https://ipecho.net/plain"

func MyIP() (net.IP, error) {
	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	ip := net.ParseIP(string(data))

	return ip, nil
}
