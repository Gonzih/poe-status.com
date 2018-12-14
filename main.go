package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/tomsteele/go-nmap"
	"gopkg.in/yaml.v2"
)

const portLimit = 996

type data struct {
	Pc   map[string]string
	Xbox map[string]string
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

type Host struct {
	Name     string
	HostName string
	nmapData *nmap.NmapRun
	nmapErr  error
}

func (h *Host) Nmap() {
	if h.nmapData != nil {
		return
	}

	rawData, err := sh("nmap", "-oX", "-", h.HostName)
	if err != nil {
		h.nmapErr = err
		return
	}
	h.nmapData, h.nmapErr = nmap.Parse(rawData)
}

func (h *Host) IsUp() (bool, error) {
	h.Nmap()

	if h.nmapErr != nil {
		return false, h.nmapErr
	}

	if len(h.nmapData.Hosts) == 0 {
		return false, fmt.Errorf("Host %s is down", h.HostName)
	}

	openPorts := 0

	for _, port := range h.nmapData.Hosts[0].Ports {
		if port.State.State == "open" {
			openPorts++
			continue
		}

		log.Printf("%d => %s", port.PortId, port.State.State)
	}

	if openPorts < portLimit {
		return false, fmt.Errorf("Host only had %d open ports", openPorts)
	}

	return true, nil
}

func main() {
	fdata, err := slurp("servers.yaml")
	must(err)
	data := &data{}
	err = yaml.Unmarshal(fdata, &data)
	must(err)

	for name, host := range data.Pc {
		h := &Host{Name: name, HostName: host}
		log.Println(name)
		log.Println(h.IsUp())
	}
}
