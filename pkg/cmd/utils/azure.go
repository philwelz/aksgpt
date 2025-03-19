/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package azure

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/fatih/color"
)

func GetSubscriptionID() string {
	// Check if the environment variable is set
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		color.Red("Environment variable ARM_SUBSCRIPTION_ID is not set.")
		os.Exit(1)
	}

	return subscriptionID
}

// Create an Azure credential
func NewAzureCLICredential() (*azidentity.AzureCLICredential, error) {
	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		color.Red("Failed to get Azure credential: %v", err)
		return nil, err
	}
	return cred, nil
}
