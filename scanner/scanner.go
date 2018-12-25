package scanner

import (
	"log"

	"gitlab.com/Gonzih/poe-status.com/sh"
)

// NmapAvailable will test if nmap is present on the system
func NmapAvailable() bool {
	out, err := sh.Sh("bash", "-c", "which nmap")
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(string(out))
	return true
}
