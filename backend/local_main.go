package main

import (
	"log"
)

func localMain() {
	data := readServersYaml()

	for name, host := range data.Pc {
		log.Println(name)
		h := &Host{Name: name, HostName: host}
		h.Scan(data.Ports)
		up, err := h.IsUp()
		h.IsHostUp = up
		log.Printf("%s is %v: %v", name, up, err)
	}
}
