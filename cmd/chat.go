/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Pradipbabar/gimini-cli/pkg"
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
			outputFormat, _ := cmd.Flags().GetString("output")

			data, err := pkg.GenerateContent(prompt)
			if err != nil {
				fmt.Printf("\033[31m[Error]\033[0m %v\n", err)
				os.Exit(1)
			}

			switch outputFormat {
			case "json":
				fmt.Printf("{\n  \"response\": %q\n}\n", data)
			default:
				fmt.Printf("\033[32m[Success]\033[0m %s\n", data)
			}
		} else {
			fmt.Println("Provide flags")
			// Add default logic here
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Flags for the chat command
	chatCmd.Flags().StringP("prompt", "p", "", "Specify the chat prompt string")
	chatCmd.Flags().StringP("save", "s", "", "Specify the file path to save the chat")
	chatCmd.Flags().StringP("output", "o", "plain", "Output format (plain|json)")

	chatCmd.MarkFlagRequired("prompt")
}
