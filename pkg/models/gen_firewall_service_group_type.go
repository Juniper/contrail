package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
// nolint
func MakeFirewallServiceGroupType() *FirewallServiceGroupType {
	return &FirewallServiceGroupType{
		//TODO(nati): Apply default

		FirewallService: MakeFirewallServiceTypeSlice(),
	}
}

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
// nolint
func InterfaceToFirewallServiceGroupType(i interface{}) *FirewallServiceGroupType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallServiceGroupType{
		//TODO(nati): Apply default

		FirewallService: InterfaceToFirewallServiceTypeSlice(m["firewall_service"]),
	}
}

// MakeFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
// nolint
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
	return []*FirewallServiceGroupType{}
}

// InterfaceToFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
// nolint
func InterfaceToFirewallServiceGroupTypeSlice(i interface{}) []*FirewallServiceGroupType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallServiceGroupType{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceGroupType(item))
	}
	return result
}
