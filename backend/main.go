package main

import (
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if os.Getenv("ENVIRONMENT") == "LAMBDA" {
		lambdaMain()
	} else {
		localMain()
	}
	duration := time.Since(start)
	log.Printf("Time spend %f seconds", duration.Seconds())
}
