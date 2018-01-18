package models

// PolicyBasedForwardingRuleType

import "encoding/json"

// PolicyBasedForwardingRuleType
type PolicyBasedForwardingRuleType struct {
	Ipv6ServiceChainAddress IpAddressType        `json:"ipv6_service_chain_address,omitempty"`
	Direction               TrafficDirectionType `json:"direction,omitempty"`
	MPLSLabel               int                  `json:"mpls_label,omitempty"`
	VlanTag                 int                  `json:"vlan_tag,omitempty"`
	SRCMac                  string               `json:"src_mac,omitempty"`
	ServiceChainAddress     string               `json:"service_chain_address,omitempty"`
	DSTMac                  string               `json:"dst_mac,omitempty"`
	Protocol                string               `json:"protocol,omitempty"`
}

// String returns json representation of the object
func (model *PolicyBasedForwardingRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleType() *PolicyBasedForwardingRuleType {
	return &PolicyBasedForwardingRuleType{
		//TODO(nati): Apply default
		ServiceChainAddress:     "",
		DSTMac:                  "",
		Protocol:                "",
		Ipv6ServiceChainAddress: MakeIpAddressType(),
		Direction:               MakeTrafficDirectionType(),
		MPLSLabel:               0,
		VlanTag:                 0,
		SRCMac:                  "",
	}
}

// MakePolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
	return []*PolicyBasedForwardingRuleType{}
}
