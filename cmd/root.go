/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/philwelz/aksgpt/cmd/inspect"

	"github.com/spf13/cobra"
)

const (
	// Repository Owner
	Owner = "philwelz"

	// Repository Name
	Repo = "aksgpt"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aksgpt",
	Short: "AKS debugging powered by AI",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(inspect.InspectCmd)
}

func init() {
	// Here you will define your flags and configuration settings.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add subcommands
	addSubCommands()
}
