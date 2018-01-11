package models

// PolicyBasedForwardingRuleType

import "encoding/json"

// PolicyBasedForwardingRuleType
type PolicyBasedForwardingRuleType struct {
	ServiceChainAddress     string               `json:"service_chain_address"`
	DSTMac                  string               `json:"dst_mac"`
	Protocol                string               `json:"protocol"`
	Ipv6ServiceChainAddress IpAddressType        `json:"ipv6_service_chain_address"`
	Direction               TrafficDirectionType `json:"direction"`
	MPLSLabel               int                  `json:"mpls_label"`
	VlanTag                 int                  `json:"vlan_tag"`
	SRCMac                  string               `json:"src_mac"`
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
		SRCMac:                  "",
		ServiceChainAddress:     "",
		DSTMac:                  "",
		Protocol:                "",
		Ipv6ServiceChainAddress: MakeIpAddressType(),
		Direction:               MakeTrafficDirectionType(),
		MPLSLabel:               0,
		VlanTag:                 0,
	}
}

// MakePolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
	return []*PolicyBasedForwardingRuleType{}
}
