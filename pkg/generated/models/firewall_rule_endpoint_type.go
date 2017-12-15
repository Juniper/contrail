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

// String returns json representation of the object
func (model *FirewallRuleEndpointType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallRuleEndpointType makes FirewallRuleEndpointType
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

// InterfaceToFirewallRuleEndpointType makes FirewallRuleEndpointType from interface
func InterfaceToFirewallRuleEndpointType(iData interface{}) *FirewallRuleEndpointType {
	data := iData.(map[string]interface{})
	return &FirewallRuleEndpointType{
		Any: data["any"].(bool),

		//{"description":"Match any workload","type":"boolean"}
		AddressGroup: data["address_group"].(string),

		//{"description":"Any workload with interface in this address-group","type":"string"}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"description":"Any workload that belongs to this subnet","type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}
		Tags: data["tags"].([]string),

		//{"description":"Any workload with tags matching tags in this list","type":"array","item":{"type":"string"}}

		TagIds: data["tag_ids"].([]int),

		//{"description":"Any workload with tags ids matching all the tags ids in this list","type":"array","item":{"type":"integer"}}
		VirtualNetwork: data["virtual_network"].(string),

		//{"description":"Any workload that belongs to this virtual network ","type":"string"}

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
