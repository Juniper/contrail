package models

// PolicyBasedForwardingRuleType

import "encoding/json"

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

func (model *PolicyBasedForwardingRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

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

func InterfaceToPolicyBasedForwardingRuleType(iData interface{}) *PolicyBasedForwardingRuleType {
	data := iData.(map[string]interface{})
	return &PolicyBasedForwardingRuleType{
		ServiceChainAddress: data["service_chain_address"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"service_chain_address","Item":null,"GoName":"ServiceChainAddress","GoType":"string"}
		DSTMac: data["dst_mac"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"dst_mac","Item":null,"GoName":"DSTMac","GoType":"string"}
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"protocol","Item":null,"GoName":"Protocol","GoType":"string"}
		Ipv6ServiceChainAddress: InterfaceToIpAddressType(data["ipv6_service_chain_address"]),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"ipv6_service_chain_address","Item":null,"GoName":"Ipv6ServiceChainAddress","GoType":"IpAddressType"}
		Direction: InterfaceToTrafficDirectionType(data["direction"]),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":"both","Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["ingress","egress","both"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TrafficDirectionType","CollectionType":"","Column":"direction","Item":null,"GoName":"Direction","GoType":"TrafficDirectionType"}
		MPLSLabel: data["mpls_label"].(int),

		//{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mpls_label","Item":null,"GoName":"MPLSLabel","GoType":"int"}
		VlanTag: data["vlan_tag"].(int),

		//{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"vlan_tag","Item":null,"GoName":"VlanTag","GoType":"int"}
		SRCMac: data["src_mac"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"src_mac","Item":null,"GoName":"SRCMac","GoType":"string"}

	}
}

func InterfaceToPolicyBasedForwardingRuleTypeSlice(data interface{}) []*PolicyBasedForwardingRuleType {
	list := data.([]interface{})
	result := MakePolicyBasedForwardingRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPolicyBasedForwardingRuleType(item))
	}
	return result
}

func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
	return []*PolicyBasedForwardingRuleType{}
}
