package main

import (
	"fmt"
	"os"

	"gitlab.com/Gonzih/poe-status.com/server"
)

var cliOptions server.Options

func init() {
	rootCmd.PersistentFlags().StringVar(&cliOptions.Host, "host", "", "host to bind to")
	rootCmd.PersistentFlags().IntVar(&cliOptions.Port, "port", 8080, "port to bind to")
	rootCmd.PersistentFlags().StringVar(&cliOptions.DatabaseURL, "database", "postgres://postgres@localhost/poestatus", "database url to use (postgres only)")
	// rootCmd.MarkFlagRequired("database")
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