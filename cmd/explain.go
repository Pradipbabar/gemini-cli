/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Pradipbabar/gimini-cli/pkg"
	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Provide an explanation",
	Long: `The explain command is used to provide detailed explanations.
It allows you to explain various aspects of your application or system.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("filepath")
		saveFilePath, _ := cmd.Flags().GetString("save")
		outputFormat, _ := cmd.Flags().GetString("output")

		// Input validation
		if filePath == "" {
			fmt.Println("\033[31m[Error]\033[0m File path is required. Use -f or --filepath to specify the file.")
			// Interactive prompt for file path
			fmt.Print("Enter file path: ")
			fmt.Scanln(&filePath)
			if filePath == "" {
				fmt.Println("\033[31m[Error]\033[0m File path is still missing. Exiting.")
				return
			}
		}

		// Validate file existence
		if _, err := pkg.SafeFileStat(filePath); err != nil {
			fmt.Printf("\033[31m[Error]\033[0m File does not exist or is not accessible: %v\n", err)
			return
		}

		var (
			data string
			err  error
		)
		if saveFilePath != "" {
			fmt.Printf("\033[36m[Info]\033[0m Explaining file: %s and saving to: %s\n", filePath, saveFilePath)
			data, err = pkg.ReadFromFileandSave(filePath, saveFilePath)
		} else {
			fmt.Printf("\033[36m[Info]\033[0m Explaining file: %s\n", filePath)
			data, err = pkg.ReadFromFile(filePath)
		}
		if err != nil {
			fmt.Printf("\033[31m[Error]\033[0m %v\n", err)
			return
		}

		// Output formatting
		switch outputFormat {
		case "json":
			fmt.Printf("{\n  \"explanation\": %q\n}\n", data)
		default:
			fmt.Println("\033[32m[Success]\033[0m Explanation:\n" + data)
		}
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

	// Flags for the explain command
	explainCmd.Flags().StringP("filepath", "f", "", "Specify the file path for explanation")
	explainCmd.Flags().StringP("save", "s", "", "Specify the file path to save the explanation")
	explainCmd.Flags().StringP("output", "o", "", "Output format: plain (default) or json")

	// Other flags and configuration settings can be added here.
}
