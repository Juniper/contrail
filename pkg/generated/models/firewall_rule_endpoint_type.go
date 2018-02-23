package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFirewallRuleEndpointType makes FirewallRuleEndpointType
func MakeFirewallRuleEndpointType() *FirewallRuleEndpointType {
	return &FirewallRuleEndpointType{
		//TODO(nati): Apply default
		AddressGroup: "",
		Subnet:       MakeSubnetType(),
		Tags:         []string{},

		TagIds: []int64{},

		VirtualNetwork: "",
		Any:            false,
	}
}

// MakeFirewallRuleEndpointType makes FirewallRuleEndpointType
func InterfaceToFirewallRuleEndpointType(i interface{}) *FirewallRuleEndpointType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallRuleEndpointType{
		//TODO(nati): Apply default
		AddressGroup: schema.InterfaceToString(m["address_group"]),
		Subnet:       InterfaceToSubnetType(m["subnet"]),
		Tags:         schema.InterfaceToStringList(m["tags"]),

		TagIds: schema.InterfaceToInt64List(m["tag_ids"]),

		VirtualNetwork: schema.InterfaceToString(m["virtual_network"]),
		Any:            schema.InterfaceToBool(m["any"]),
	}
}

// MakeFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
	return []*FirewallRuleEndpointType{}
}

// InterfaceToFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
func InterfaceToFirewallRuleEndpointTypeSlice(i interface{}) []*FirewallRuleEndpointType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallRuleEndpointType{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleEndpointType(item))
	}
	return result
}
