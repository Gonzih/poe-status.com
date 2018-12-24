package main

import (
	"log"

	"gitlab.com/Gonzih/poe-status.com/server"
)

func main() {
	port := 8080
	log.Printf("Starting server on port %d", port)
	log.Fatal(server.StartServer(port))
}
