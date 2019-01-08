package main

import (
	"github.com/spf13/cobra"
	"github.com/Gonzih/poe-status.com/app/server"
)

var rootCmd = &cobra.Command{
	Use:   "poe-status-server",
	Short: "Command line interface for poe-status server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.StartServer(&cliOptions)
	},
}
