package main

import (
	"fmt"
	"os"

	"gitlab.com/Gonzih/poe-status.com/client"
)

var cliOptions client.Options

func init() {
	rootCmd.PersistentFlags().StringVar(&cliOptions.HostPort, "host", "localhost:8080", "poe-status instance host and port in localhost:8080 format")
	rootCmd.MarkFlagRequired("host")
}

func main() {
	execute()
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
