/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Initiate a chat session",
	Long: `The chat command is used to initiate a chat session.
It allows you to interact with the chat functionality of your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt, _ := cmd.Flags().GetString("prompt")
		saveFilePath, _ := cmd.Flags().GetString("save")

		if prompt != "" && saveFilePath != "" {
			fmt.Printf("Chat called with prompt: %s and save file path: %s\n", prompt, saveFilePath)
			// Add logic for when both -p and -s are present
		} else if prompt != "" {
			fmt.Printf("Chat called with prompt: %s\n", prompt)
			// Add logic for when only -p is present
		} else {
			fmt.Println("Chat called")
			// Add default logic here
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Flags for the chat command
	chatCmd.Flags().StringP("prompt", "p", "", "Specify the chat prompt string")
	chatCmd.Flags().StringP("save", "s", "", "Specify the file path to save the chat")

	// Other flags and configuration settings can be added here.
}
