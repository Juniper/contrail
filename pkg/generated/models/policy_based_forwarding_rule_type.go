package models

// PolicyBasedForwardingRuleType

import "encoding/json"

// PolicyBasedForwardingRuleType
type PolicyBasedForwardingRuleType struct {
	Ipv6ServiceChainAddress IpAddressType        `json:"ipv6_service_chain_address"`
	Direction               TrafficDirectionType `json:"direction"`
	MPLSLabel               int                  `json:"mpls_label"`
	VlanTag                 int                  `json:"vlan_tag"`
	SRCMac                  string               `json:"src_mac"`
	ServiceChainAddress     string               `json:"service_chain_address"`
	DSTMac                  string               `json:"dst_mac"`
	Protocol                string               `json:"protocol"`
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
		Ipv6ServiceChainAddress: MakeIpAddressType(),
		Direction:               MakeTrafficDirectionType(),
		MPLSLabel:               0,
		VlanTag:                 0,
		SRCMac:                  "",
		ServiceChainAddress:     "",
		DSTMac:                  "",
		Protocol:                "",
	}
}

// InterfaceToPolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType from interface
func InterfaceToPolicyBasedForwardingRuleType(iData interface{}) *PolicyBasedForwardingRuleType {
	data := iData.(map[string]interface{})
	return &PolicyBasedForwardingRuleType{
		SRCMac: data["src_mac"].(string),

		//{"type":"string"}
		ServiceChainAddress: data["service_chain_address"].(string),

		//{"type":"string"}
		DSTMac: data["dst_mac"].(string),

		//{"type":"string"}
		Protocol: data["protocol"].(string),

		//{"type":"string"}
		Ipv6ServiceChainAddress: InterfaceToIpAddressType(data["ipv6_service_chain_address"]),

		//{"type":"string"}
		Direction: InterfaceToTrafficDirectionType(data["direction"]),

		//{"default":"both","type":"string","enum":["ingress","egress","both"]}
		MPLSLabel: data["mpls_label"].(int),

		//{"type":"integer"}
		VlanTag: data["vlan_tag"].(int),

		//{"type":"integer"}

	}
}

// InterfaceToPolicyBasedForwardingRuleTypeSlice makes a slice of PolicyBasedForwardingRuleType from interface
func InterfaceToPolicyBasedForwardingRuleTypeSlice(data interface{}) []*PolicyBasedForwardingRuleType {
	list := data.([]interface{})
	result := MakePolicyBasedForwardingRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPolicyBasedForwardingRuleType(item))
	}
	return result
}

// MakePolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
	return []*PolicyBasedForwardingRuleType{}
}
