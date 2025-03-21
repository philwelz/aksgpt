/*
Copyright © 2025 Philip Welz
*/
package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/fatih/color"

	"github.com/philwelz/aksgpt/pkg/azure"
)

// ClusterInfo represents the AKS cluster information structure
type ClusterInfo struct {
	// Name              string      `json:"name"`
	// ResourceGroup     string      `json:"resourceGroup"`
	// Location          string      `json:"location"`
	// KubernetesVersion string      `json:"kubernetesVersion"`
	// ProvisioningState string      `json:"provisioningState"`
	// NodeResourceGroup string      `json:"nodeResourceGroup"`
	// FQDN              string      `json:"fqdn"`
	Properties interface{} `json:"properties"`
}

func GetClusterInfo(subscriptionID, resourceGroup, clusterName string) (interface{}, error) {
	// Call Azure authentication function
	cred, err := azure.NewAzureCLICredential()
	if err != nil {
		color.Red("Failed to create Azure credential: %v", err)
		return nil, err
	}

	// Create an AKS client
	client, err := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)
	if err != nil {
		color.Red("Failed to create AKS client: %v", err)
		return nil, err
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Fetch cluster information
	cluster, err := client.Get(ctx, resourceGroup, clusterName, nil)
	if err != nil {
		color.Red("Failed to get AKS cluster %s in Subscription %s", clusterName, subscriptionID)
		return nil, err
	}

	// Validate cluster properties
	if cluster.Properties == nil {
		color.Red("Error: cluster properties are nil")
		return nil, fmt.Errorf("cluster properties are nil")
	}

	// Convert the cluster properties to a map
	propBytes, err := json.Marshal(cluster.Properties)
	if err != nil {
		color.Red("Failed to marshal cluster properties: %v", err)
		return nil, err
	}

	// Create a map to hold the properties
	var propsMap map[string]interface{}

	// Unmarshal the properties into a map
	if unmarshalErr := json.Unmarshal(propBytes, &propsMap); unmarshalErr != nil {
		color.Red("Failed to unmarshal cluster properties: %v", unmarshalErr)
		return nil, unmarshalErr
	}
	// Anonymize sensitive fields in properties
	anonymizedProps := anonymizeSensitiveFields(propsMap)

	// Create cluster info struct with anonymized properties
	clusterInfo := ClusterInfo{
		Properties: anonymizedProps,
	}

	// Convert to JSON with proper indentation
	output, err := json.MarshalIndent(clusterInfo, "", "  ")
	if err != nil {
		color.Red("Failed to marshal cluster info: %v", err)
		return nil, err
	}

	// Print the cluster information
	return string(output), nil
}
