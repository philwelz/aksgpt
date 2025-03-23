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
	Short: "Display the aksgpt version number.",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("aksgpt: %s\n", "v0.0.3")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
