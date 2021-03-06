package main

import (
	"github.com/spf13/cobra"
	"github.com/Gonzih/poe-status.com/app/client"
)

var rootCmd = &cobra.Command{
	Use:   "poe-status-cli",
	Short: "Command line interface for poe-status reporting",
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Call(&cliOptions)
	},
}
