package models

// FirewallRuleEndpointType

import "encoding/json"

// FirewallRuleEndpointType
type FirewallRuleEndpointType struct {
	Subnet         *SubnetType `json:"subnet"`
	Tags           []string    `json:"tags"`
	TagIds         []int       `json:"tag_ids"`
	VirtualNetwork string      `json:"virtual_network"`
	Any            bool        `json:"any"`
	AddressGroup   string      `json:"address_group"`
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
		VirtualNetwork: "",
		Any:            false,
		AddressGroup:   "",
		Subnet:         MakeSubnetType(),
		Tags:           []string{},

		TagIds: []int{},
	}
}

// MakeFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
	return []*FirewallRuleEndpointType{}
}
