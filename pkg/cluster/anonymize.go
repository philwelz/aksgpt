/*
Copyright © 2025 Philip Welz
*/
package cluster

// anonymizeMap recursively processes a map to anonymize sensitive fields
func anonymizeSensitiveFields(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			// Define sensitive fields that need to be redacted
			sensitiveFields := map[string]bool{
				"clientId":             true,
				"objectId":             true,
				"resourceId":           true,
				"tenantID":             true,
				"azurePortalFQDN":      true,
				"dnsPrefix":            true,
				"fqdn":                 true,
				"effectiveOutboundIPs": true,
				"nodeResourceGroup":    true,
				"location":             true,
				"subscriptionId":       true,
				"adminUsername":        true,
				"podSubnetID":          true,
				"vnetSubnetID":         true,
			}

			if sensitiveFields[key] {
				v[key] = "REDACTED"
			} else {
				v[key] = anonymizeSensitiveFields(val)
			}
		}
		return v
	case []interface{}:
		for i, val := range v {
			v[i] = anonymizeSensitiveFields(val)
		}
		return v
	default:
		return v
	}
}
