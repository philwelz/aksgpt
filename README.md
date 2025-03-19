
`aksgpt` is a [k8sgpt](https://github.com/k8sgpt-ai/k8sgpt) inspired tool leveraging AI for scanning your AKS cluster configuration and providing recommendations for best practices. It is designed to help you identify potential issues and improve the security and performance of your AKS clusters.

## Installation

To install `aksgpt`, you can use the following command:

```bash
go install github.com/philwelz/aksgpt
```

## AI Backends

`aksgpt` supports the following AI backends to provide recommendations:
- [OpenAI](https://platform.openai.com/)

### Quick Start

To run `aksgpt`, you need to set up the following environment variables:

```bash
export ARM_SUBSCRIPTION_ID=<Azure Subscription ID>
export OPENAI_API_KEY=<Platform API Key>
```

You can then run `aksgpt` with the following command:

```bash
aksgpt inspect cluster -g <Resource Group> -c <Cluster Name>
```

IMPORTANT: Remember to add $GOPATH to your $PATH variable. Otherwise, you find the binary in $GOPATH/bin.

## Anonymized cluster information

The cluster information is anonymized before sending it to the AI backend. The following fields are **REDACTED** as they could contain sensitive information:

- `tenantID`
- `subscriptionId`
- `clientId`
- `objectId`
- `resourceId`
- `azurePortalFQDN`
- `fqdn`
- `dnsPrefix`
- `effectiveOutboundIPs`
- `nodeResourceGroup`
- `location`

Sample AKS config json that is send to the AI backend:

<details>
<summary>JSON</summary>

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




