package models

// VrfAssignRuleType

import "encoding/json"

// VrfAssignRuleType
type VrfAssignRuleType struct {
	RoutingInstance string              `json:"routing_instance"`
	MatchCondition  *MatchConditionType `json:"match_condition"`
	VlanTag         int                 `json:"vlan_tag"`
	IgnoreACL       bool                `json:"ignore_acl"`
}

//  parents relation object

// String returns json representation of the object
func (model *VrfAssignRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVrfAssignRuleType makes VrfAssignRuleType
func MakeVrfAssignRuleType() *VrfAssignRuleType {
	return &VrfAssignRuleType{
		//TODO(nati): Apply default
		RoutingInstance: "",
		MatchCondition:  MakeMatchConditionType(),
		VlanTag:         0,
		IgnoreACL:       false,
	}
}

// InterfaceToVrfAssignRuleType makes VrfAssignRuleType from interface
func InterfaceToVrfAssignRuleType(iData interface{}) *VrfAssignRuleType {
	data := iData.(map[string]interface{})
	return &VrfAssignRuleType{
		MatchCondition: InterfaceToMatchConditionType(data["match_condition"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dst_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string","GoPremitive":true},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string","GoPremitive":true},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType","GoPremitive":false},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType","GoPremitive":false},"GoName":"SubnetList","GoType":"[]*SubnetType","GoPremitive":true},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"DSTAddress","GoType":"AddressType","GoPremitive":false},"dst_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType","GoPremitive":false},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPort","GoType":"PortType","GoPremitive":false},"ethertype":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["IPv4","IPv6"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/EtherType","CollectionType":"","Column":"","Item":null,"GoName":"Ethertype","GoType":"EtherType","GoPremitive":false},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true},"src_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"network_policy":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string","GoPremitive":true},"security_group":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string","GoPremitive":true},"subnet":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType","GoPremitive":false},"subnet_list":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType","GoPremitive":false},"GoName":"SubnetList","GoType":"[]*SubnetType","GoPremitive":true},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressType","CollectionType":"","Column":"","Item":null,"GoName":"SRCAddress","GoType":"AddressType","GoPremitive":false},"src_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType","GoPremitive":false},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPort","GoType":"PortType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MatchConditionType","CollectionType":"","Column":"","Item":null,"GoName":"MatchCondition","GoType":"MatchConditionType","GoPremitive":false}
		VlanTag: data["vlan_tag"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VlanTag","GoType":"int","GoPremitive":true}
		IgnoreACL: data["ignore_acl"].(bool),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IgnoreACL","GoType":"bool","GoPremitive":true}
		RoutingInstance: data["routing_instance"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToVrfAssignRuleTypeSlice makes a slice of VrfAssignRuleType from interface
func InterfaceToVrfAssignRuleTypeSlice(data interface{}) []*VrfAssignRuleType {
	list := data.([]interface{})
	result := MakeVrfAssignRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVrfAssignRuleType(item))
	}
	return result
}

// MakeVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
	return []*VrfAssignRuleType{}
}
