package main

import (
	"fmt"
	"os"

	"gitlab.com/Gonzih/poe-status.com/app/client"
)

var cliOptions client.Options

func init() {
	rootCmd.PersistentFlags().StringVar(&cliOptions.URL, "url", "http://localhost:8080", "poe-status instance url")
	rootCmd.MarkFlagRequired("url")
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
