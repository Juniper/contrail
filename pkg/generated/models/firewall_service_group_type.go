package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
func MakeFirewallServiceGroupType() *FirewallServiceGroupType {
	return &FirewallServiceGroupType{
		//TODO(nati): Apply default

		FirewallService: MakeFirewallServiceTypeSlice(),
	}
}

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
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
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
	return []*FirewallServiceGroupType{}
}

// InterfaceToFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
func InterfaceToFirewallServiceGroupTypeSlice(i interface{}) []*FirewallServiceGroupType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallServiceGroupType{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceGroupType(item))
	}
	return result
}
