package models

// FirewallRuleEndpointType

import "encoding/json"

// FirewallRuleEndpointType
type FirewallRuleEndpointType struct {
	Tags           []string    `json:"tags,omitempty"`
	TagIds         []int       `json:"tag_ids,omitempty"`
	VirtualNetwork string      `json:"virtual_network,omitempty"`
	Any            bool        `json:"any"`
	AddressGroup   string      `json:"address_group,omitempty"`
	Subnet         *SubnetType `json:"subnet,omitempty"`
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
		AddressGroup: "",
		Subnet:       MakeSubnetType(),
		Tags:         []string{},

		TagIds: []int{},

		VirtualNetwork: "",
		Any:            false,
	}
}

// MakeFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
	return []*FirewallRuleEndpointType{}
}
