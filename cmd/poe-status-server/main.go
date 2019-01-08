package main

import (
	"fmt"
	"os"

	"github.com/Gonzih/poe-status.com/app/server"
)

var cliOptions server.Options

func init() {
	rootCmd.PersistentFlags().StringVar(&cliOptions.Host, "host", "", "host to bind to")
	rootCmd.PersistentFlags().IntVar(&cliOptions.Port, "port", 8080, "port to bind to")
	rootCmd.PersistentFlags().StringVar(&cliOptions.DatabaseURL, "database", "postgres://postgres@localhost/poestatus", "database url to use (postgres only)")
	rootCmd.PersistentFlags().StringVar(&cliOptions.MigrationsFolderPath, "migrations-folder", "./migrations", "path to migrations folder")
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
