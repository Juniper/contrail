package models

// PolicyBasedForwardingRuleType

import "encoding/json"

// PolicyBasedForwardingRuleType
type PolicyBasedForwardingRuleType struct {
	SRCMac                  string               `json:"src_mac"`
	ServiceChainAddress     string               `json:"service_chain_address"`
	DSTMac                  string               `json:"dst_mac"`
	Protocol                string               `json:"protocol"`
	Ipv6ServiceChainAddress IpAddressType        `json:"ipv6_service_chain_address"`
	Direction               TrafficDirectionType `json:"direction"`
	MPLSLabel               int                  `json:"mpls_label"`
	VlanTag                 int                  `json:"vlan_tag"`
}

//  parents relation object

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

// InterfaceToPolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType from interface
func InterfaceToPolicyBasedForwardingRuleType(iData interface{}) *PolicyBasedForwardingRuleType {
	data := iData.(map[string]interface{})
	return &PolicyBasedForwardingRuleType{
		ServiceChainAddress: data["service_chain_address"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"service_chain_address","Item":null,"GoName":"ServiceChainAddress","GoType":"string","GoPremitive":true}
		DSTMac: data["dst_mac"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"dst_mac","Item":null,"GoName":"DSTMac","GoType":"string","GoPremitive":true}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"protocol","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true}
		Ipv6ServiceChainAddress: InterfaceToIpAddressType(data["ipv6_service_chain_address"]),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"ipv6_service_chain_address","Item":null,"GoName":"Ipv6ServiceChainAddress","GoType":"IpAddressType","GoPremitive":false}
		Direction: InterfaceToTrafficDirectionType(data["direction"]),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":"both","Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["ingress","egress","both"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TrafficDirectionType","CollectionType":"","Column":"direction","Item":null,"GoName":"Direction","GoType":"TrafficDirectionType","GoPremitive":false}
		MPLSLabel: data["mpls_label"].(int),

		//{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mpls_label","Item":null,"GoName":"MPLSLabel","GoType":"int","GoPremitive":true}
		VlanTag: data["vlan_tag"].(int),

		//{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"vlan_tag","Item":null,"GoName":"VlanTag","GoType":"int","GoPremitive":true}
		SRCMac: data["src_mac"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"src_mac","Item":null,"GoName":"SRCMac","GoType":"string","GoPremitive":true}

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
