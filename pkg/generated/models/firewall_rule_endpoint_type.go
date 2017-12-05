package models

// FirewallRuleEndpointType

import "encoding/json"

type FirewallRuleEndpointType struct {
	TagIds         []int       `json:"tag_ids"`
	VirtualNetwork string      `json:"virtual_network"`
	Any            bool        `json:"any"`
	AddressGroup   string      `json:"address_group"`
	Subnet         *SubnetType `json:"subnet"`
	Tags           []string    `json:"tags"`
}

func (model *FirewallRuleEndpointType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFirewallRuleEndpointType() *FirewallRuleEndpointType {
	return &FirewallRuleEndpointType{
		//TODO(nati): Apply default
		Subnet: MakeSubnetType(),
		Tags:   []string{},

		TagIds: []int{},

		VirtualNetwork: "",
		Any:            false,
		AddressGroup:   "",
	}
}

func InterfaceToFirewallRuleEndpointType(iData interface{}) *FirewallRuleEndpointType {
	data := iData.(map[string]interface{})
	return &FirewallRuleEndpointType{
		VirtualNetwork: data["virtual_network"].(string),

		//{"Title":"","Description":"Any workload that belongs to this virtual network ","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}
		Any: data["any"].(bool),

		//{"Title":"","Description":"Match any workload","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Any","GoType":"bool"}
		AddressGroup: data["address_group"].(string),

		//{"Title":"","Description":"Any workload with interface in this address-group","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AddressGroup","GoType":"string"}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"Title":"","Description":"Any workload that belongs to this subnet","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"}
		Tags: data["tags"].([]string),

		//{"Title":"","Description":"Any workload with tags matching tags in this list","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tags","GoType":"string"},"GoName":"Tags","GoType":"[]string"}

		TagIds: data["tag_ids"].([]int),

		//{"Title":"","Description":"Any workload with tags ids matching all the tags ids in this list","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagIds","GoType":"int"},"GoName":"TagIds","GoType":"[]int"}

	}
}

func InterfaceToFirewallRuleEndpointTypeSlice(data interface{}) []*FirewallRuleEndpointType {
	list := data.([]interface{})
	result := MakeFirewallRuleEndpointTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleEndpointType(item))
	}
	return result
}

func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
	return []*FirewallRuleEndpointType{}
}
