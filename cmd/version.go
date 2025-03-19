/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number aksgpt",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("aksgpt: %s\n", "0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
