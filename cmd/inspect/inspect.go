/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package inspect

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// InspectCmd represents the inspect command
var InspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Analyze your AKS cluster to ensure compliance with best practices.",
	Long:  ``,
	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			color.Red("Error displaying help:", err)
		}
	},
}

// ClusterName and ResourceGroup are global variables to hold the cluster name and resource group
var (
	ClusterName   string
	ResourceGroup string
	Backend       string
)

func init() {
	// Define the global flags for the inspect command
	InspectCmd.PersistentFlags().
		StringVarP(&ClusterName, "cluster-name", "c", "", "The name of the AKS cluster")
	InspectCmd.PersistentFlags().
		StringVarP(&ResourceGroup, "resource-group", "g", "", "The name of the Resource Group")
	InspectCmd.PersistentFlags().
		StringVarP(&Backend, "backend", "b", "openai", "The AI backend to use (openai, azure, defaults to openai)")

	// Mark flags as required
	if err := InspectCmd.MarkPersistentFlagRequired("cluster-name"); err != nil {
		panic(err)
	}
	if err := InspectCmd.MarkPersistentFlagRequired("resource-group"); err != nil {
		panic(err)
	}
}
