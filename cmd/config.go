/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long: `The config command is used to manage configuration settings for your application.
It allows you to view, set, and modify various configuration options.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		fmt.Printf("config called with key: %s\n", key)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Flags for the config command
	configCmd.Flags().StringP("key", "k", "", "Specify the configuration key (required)")

	// Other flags and configuration settings can be added here.
}
