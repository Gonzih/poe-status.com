package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")

	router := httprouter.New()
	router.GET("/status", StatusHandler)

	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf(":%s", port)

	log.Fatal(http.ListenAndServe(address, router))
}
