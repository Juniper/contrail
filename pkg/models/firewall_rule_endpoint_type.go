package models

import "github.com/Juniper/contrail/pkg/common"

// ValidateEndpointType checks if endpoint refers to only one endpoint type
func (e *FirewallRuleEndpointType) ValidateEndpointType() error {
	if e == nil {
		return nil
	}

	count := 0
	if e.GetAddressGroup() != "" {
		count++
	}
	if e.GetAny() == true {
		count++
	}
	if e.GetSubnet() != nil {
		count++
	}
	if len(e.GetTags()) > 0 {
		count++
	}
	if e.GetVirtualNetwork() != "" {
		count++
	}

	if count > 1 {
		return common.ErrorBadRequest("endpoint is limited to only one endpoint type at a time")
	}

	return nil
}
