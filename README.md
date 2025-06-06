[![CodeQL](https://github.com/philwelz/aksgpt/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/philwelz/aksgpt/actions/workflows/github-code-scanning/codeql)
[![ci](https://github.com/philwelz/aksgpt/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/philwelz/aksgpt/actions/workflows/golangci-lint.yaml)


![GitHub last commit (branch)](https://img.shields.io/github/last-commit/philwelz/aksgpt/main)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/philwelz/aksgpt)

[![Go Report Card](https://goreportcard.com/badge/github.com/philwelz/aksgpt)](https://goreportcard.com/report/github.com/philwelz/aksgpt)
[![Go version](https://img.shields.io/github/go-mod/go-version/philwelz/aksgpt.svg)](https://github.com/philwelz/aksgpt)


`aksgpt` is a powerful tool inspired by [k8sgpt](https://github.com/k8sgpt-ai/k8sgpt)  that uses AI to analyze and review your Azure Kubernetes Service (AKS) cluster configuration. It provides actionable recommendations for best practices, helping you identify potential issues and optimize the security, performance, and overall health of your AKS clusters.

## Installation

To install `aksgpt`, you can use the following command:

<!---x-release-please-start-version-->
```bash
go install github.com/philwelz/aksgpt@v0.0.5
```
<!---x-release-please-end-->

> IMPORTANT: Make sure to add $GOPATH to your $PATH variable. If you don't, the binary will be located in $GOPATH/bin.

## AI Backends

`aksgpt` integrates with the following AI backends to deliver recommendations:

- [GitHub Models](https://github.com/marketplace/models) - the default backend if no other is specified
- [OpenAI](https://platform.openai.com/)
- [Azure OpenAI](https://azure.microsoft.com/products/ai-services/openai-service)


### Quick Start

Before running `aksgpt` with the default **GitHub Models** backend, ensure that the following environment variables are configured:

```bash
export ARM_SUBSCRIPTION_ID=<Azure Subscription ID>
export GITHUB_TOKEN=<GitHub fine-grained PAT>
```

If you are using **Azure OpenAI**, set the following environment variables:

```bash
export ARM_SUBSCRIPTION_ID=<Azure Subscription ID>
export AZURE_OPENAI_API_KEY=<Azure Open AI API Key>
export AZURE_OPENAI_ENDPOINT=<Azure OpenAI Endpoint>
```

If you are using **OpenAi**, set the following environment variables:

```bash
export ARM_SUBSCRIPTION_ID=<Azure Subscription ID>
export OPENAI_API_KEY=<Platform API Key>
```

You can run `aksgpt` using the command below:

```bash
aksgpt inspect cluster -g <Resource Group> -c <Cluster Name> -b <Backend>
```

## Anonymized cluster information

### Redacted Fields

Cluster information is anonymized prior to being sent to the AI backend. The following fields are REDACTED to ensure sensitive data is protected:

<details>

  <summary>List of redacted Fields</summary>

- adminGroupObjectIDs
- adminUsername
- azurePortalFQDN
- clientId
- dnsPrefix
- effectiveOutboundIPs
- fqdn
- id
- issuerURL
- location
- nodeResourceGroup
- objectId
- podSubnetID
- principalId
- privateDnsZone
- privateFQDN
- resourceId
- subscriptionId
- tenantID
- userAssignedIdentities
- vnetSubnetID
</details>


### Example AKS configuration

The following is an example of the AKS configuration JSON that is submitted to the AI backend.

<details>

  <summary>Example AKS configuration JSON submitted to the AI backend</summary>

```json
{
  "properties": {
    "aadProfile": {
      "enableAzureRBAC": true,
      "managed": true,
      "tenantID": "REDACTED"
    },
    "addonProfiles": {
      "azurepolicy": {
        "config": {
          "version": "v2"
        },
        "enabled": true,
        "identity": {
          "clientId": "REDACTED",
          "objectId": "REDACTED",
          "resourceId": "REDACTED"
        }
      },
      "extensionManager": {
        "enabled": true,
        "identity": {
          "clientId": "REDACTED",
          "objectId": "REDACTED",
          "resourceId": "REDACTED"
        }
      }
    },
    "agentPoolProfiles": [
      {
        "availabilityZones": [
          "2",
          "3",
          "1"
        ],
        "count": 1,
        "currentOrchestratorVersion": "1.31.2",
        "enableAutoScaling": true,
        "enableEncryptionAtHost": false,
        "enableFIPS": false,
        "enableNodePublicIP": false,
        "enableUltraSSD": false,
        "kubeletDiskType": "OS",
        "maxCount": 2,
        "maxPods": 250,
        "minCount": 1,
        "mode": "System",
        "name": "sy0aksbackup",
        "nodeImageVersion": "AKSUbuntu-2204gen2containerd-202411.12.0",
        "orchestratorVersion": "1.31.2",
        "osDiskSizeGB": 128,
        "osDiskType": "Managed",
        "osSKU": "Ubuntu",
        "osType": "Linux",
        "powerState": {
          "code": "Stopped"
        },
        "provisioningState": "Succeeded",
        "scaleDownMode": "Delete",
        "tags": {
          "location": "REDACTED",
          "maintainer": "029b5baf-7254-4d02-8b55-7d43a0a04926",
          "project": "aksbackup",
          "stage": "pet"
        },
        "type": "VirtualMachineScaleSets",
        "upgradeSettings": {
          "maxSurge": "33%"
        },
        "vmSize": "Standard_D2ds_v5"
      }
    ],
    "apiServerAccessProfile": {
      "disableRunCommand": true
    },
    "autoScalerProfile": {
      "balance-similar-node-groups": "false",
      "expander": "random",
      "max-empty-bulk-delete": "10",
      "max-graceful-termination-sec": "600",
      "max-node-provision-time": "15m",
      "max-total-unready-percentage": "45",
      "new-pod-scale-up-delay": "0s",
      "ok-total-unready-count": "3",
      "scale-down-delay-after-add": "10m",
      "scale-down-delay-after-delete": "10s",
      "scale-down-delay-after-failure": "3m",
      "scale-down-unneeded-time": "10m",
      "scale-down-unready-time": "20m",
      "scale-down-utilization-threshold": "0.5",
      "scan-interval": "10s",
      "skip-nodes-with-local-storage": "false",
      "skip-nodes-with-system-pods": "true"
    },
    "autoUpgradeProfile": {
      "upgradeChannel": "none"
    },
    "azurePortalFQDN": "REDACTED",
    "currentKubernetesVersion": "1.31.2",
    "disableLocalAccounts": true,
    "dnsPrefix": "REDACTED",
    "enableRBAC": true,
    "fqdn": "REDACTED",
    "identityProfile": {
      "kubeletidentity": {
        "clientId": "REDACTED",
        "objectId": "REDACTED",
        "resourceId": "REDACTED"
      }
    },
    "kubernetesVersion": "1.31.2",
    "maxAgentPools": 100,
    "networkProfile": {
      "dnsServiceIP": "10.0.0.10",
      "ipFamilies": [
        "IPv4"
      ],
      "loadBalancerProfile": {
        "effectiveOutboundIPs": "REDACTED",
        "managedOutboundIPs": {
          "count": 1
        }
      },
      "loadBalancerSku": "standard",
      "networkPlugin": "azure",
      "networkPolicy": "azure",
      "outboundType": "loadBalancer",
      "serviceCidr": "10.0.0.0/16",
      "serviceCidrs": [
        "10.0.0.0/16"
      ]
    },
    "nodeResourceGroup": "REDACTED",
    "powerState": {
      "code": "Stopped"
    },
    "provisioningState": "Succeeded",
    "securityProfile": {},
    "servicePrincipalProfile": {
      "clientId": "REDACTED"
    },
    "storageProfile": {
      "diskCSIDriver": {
        "enabled": true
      },
      "fileCSIDriver": {
        "enabled": true
      },
      "snapshotController": {
        "enabled": true
      }
    },
    "windowsProfile": {
      "adminUsername": "REDACTED",
      "enableCSIProxy": true
    }
  }
}
```
</details>
