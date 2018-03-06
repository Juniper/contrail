package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallRuleEndpointType makes FirewallRuleEndpointType
// nolint
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
// nolint
func InterfaceToFirewallRuleEndpointType(i interface{}) *FirewallRuleEndpointType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallRuleEndpointType{
		//TODO(nati): Apply default
		AddressGroup: common.InterfaceToString(m["address_group"]),
		Subnet:       InterfaceToSubnetType(m["subnet"]),
		Tags:         common.InterfaceToStringList(m["tags"]),

		TagIds: common.InterfaceToInt64List(m["tag_ids"]),

		VirtualNetwork: common.InterfaceToString(m["virtual_network"]),
		Any:            common.InterfaceToBool(m["any"]),
	}
}

// MakeFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
// nolint
func MakeFirewallRuleEndpointTypeSlice() []*FirewallRuleEndpointType {
	return []*FirewallRuleEndpointType{}
}

// InterfaceToFirewallRuleEndpointTypeSlice() makes a slice of FirewallRuleEndpointType
// nolint
func InterfaceToFirewallRuleEndpointTypeSlice(i interface{}) []*FirewallRuleEndpointType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallRuleEndpointType{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleEndpointType(item))
	}
	return result
}
