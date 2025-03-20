/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package inspect

import (
	ai "github.com/philwelz/aksgpt/pkg/ai"
	cluster "github.com/philwelz/aksgpt/pkg/cmd/inspect"
	azure "github.com/philwelz/aksgpt/pkg/cmd/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// clusterCmd represents the openai command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Analyze your AKS cluster to ensure compliance with best practices.",
	Long:  ``,
	Run: func(cmd *cobra.Command, _ []string) {
		// Check if required flags are provided
		if ClusterName == "" || ResourceGroup == "" {
			color.Red("Error: cluster-name and resource-group are required")
			if err := cmd.Help(); err != nil {
				color.Red("Error displaying help:", err)
			}
			return
		}

		// use globals from azure pkg
		subscriptionID := azure.GetSubscriptionID()

		// use Globals from inspect.go
		resourceGroup := ResourceGroup
		clusterName := ClusterName

		// Run the cluster inspection function
		info, err := cluster.GetClusterInfo(subscriptionID, resourceGroup, clusterName)
		if err != nil {
			color.Red("Error getting cluster info: %v\n", err)
			return
		}

		// Display the cluster information - usedful for debugging
		// fmt.Println(info.(string))

		ai.OpenAiChat(info.(string))
	},
}

func init() {
	// Add clusterCmd to the inspect package
	InspectCmd.AddCommand(clusterCmd)
}
