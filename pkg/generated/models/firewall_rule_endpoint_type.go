package models

// FirewallRuleEndpointType

import "encoding/json"

// FirewallRuleEndpointType
type FirewallRuleEndpointType struct {
	AddressGroup   string      `json:"address_group"`
	Subnet         *SubnetType `json:"subnet"`
	Tags           []string    `json:"tags"`
	TagIds         []int       `json:"tag_ids"`
	VirtualNetwork string      `json:"virtual_network"`
	Any            bool        `json:"any"`
}

//  parents relation object

// String returns json representation of the object
func (model *FirewallRuleEndpointType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallRuleEndpointType makes FirewallRuleEndpointType
func MakeFirewallRuleEndpointType() *FirewallRuleEndpointType {
	return &FirewallRuleEndpointType{
		//TODO(nati): Apply default

		TagIds: []int{},

		VirtualNetwork: "",
		Any:            false,
		AddressGroup:   "",
		Subnet:         MakeSubnetType(),
		Tags:           []string{},
	}
}

// InterfaceToFirewallRuleEndpointType makes FirewallRuleEndpointType from interface
func InterfaceToFirewallRuleEndpointType(iData interface{}) *FirewallRuleEndpointType {
	data := iData.(map[string]interface{})
	return &FirewallRuleEndpointType{
		Any: data["any"].(bool),

		//{"Title":"","Description":"Match any workload","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Any","GoType":"bool","GoPremitive":true}
		AddressGroup: data["address_group"].(string),

		//{"Title":"","Description":"Any workload with interface in this address-group","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AddressGroup","GoType":"string","GoPremitive":true}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"Title":"","Description":"Any workload that belongs to this subnet","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType","GoPremitive":false}
		Tags: data["tags"].([]string),

		//{"Title":"","Description":"Any workload with tags matching tags in this list","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tags","GoType":"string","GoPremitive":true},"GoName":"Tags","GoType":"[]string","GoPremitive":true}

		TagIds: data["tag_ids"].([]int),

		//{"Title":"","Description":"Any workload with tags ids matching all the tags ids in this list","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagIds","GoType":"int","GoPremitive":true},"GoName":"TagIds","GoType":"[]int","GoPremitive":true}
		VirtualNetwork: data["virtual_network"].(string),

		//{"Title":"","Description":"Any workload that belongs to this virtual network ","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToFirewallRuleEndpointTypeSlice makes a slice of FirewallRuleEndpointType from interface
func InterfaceToFirewallRuleEndpointTypeSlice(data interface{}) []*FirewallRuleEndpointType {
	list := data.([]interface{})
	result := MakeFirewallRuleEndpointTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleEndpointType(item))
	}
	return result
}

// MakeFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
	return []*FirewallRuleEndpointType{}
}
