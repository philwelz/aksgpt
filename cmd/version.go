/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const PackageVersion = "v0.0.3" // x-release-please-version

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the aksgpt version number.",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("aksgpt: %s\n", PackageVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
