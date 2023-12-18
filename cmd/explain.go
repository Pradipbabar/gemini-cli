/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Pradipbabar/gimini-cli/pkg"
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

		if filePath != "" && saveFilePath != "" {
			fmt.Printf("Explain called with file path: %s and save file path: %s\n", filePath, saveFilePath)
			data, err := pkg.ReadFromFileandSave(filePath,saveFilePath)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data)
		
		} else if filePath != "" {
			fmt.Printf("Explain called with file path: %s\n", filePath)
			data, err := pkg.ReadFromFile(filePath)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data)
		} else {
			fmt.Println("Provide flags")
			// Add default logic here
		}
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

	// Flags for the explain command
	explainCmd.Flags().StringP("filepath", "f", "", "Specify the file path for explanation")
	explainCmd.Flags().StringP("save", "s", "", "Specify the file path to save the explanation")

	// Other flags and configuration settings can be added here.
}
