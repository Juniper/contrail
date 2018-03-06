package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFirewallServiceType makes FirewallServiceType
// nolint
func MakeFirewallServiceType() *FirewallServiceType {
	return &FirewallServiceType{
		//TODO(nati): Apply default
		Protocol:   "",
		DSTPorts:   MakePortType(),
		SRCPorts:   MakePortType(),
		ProtocolID: 0,
	}
}

// MakeFirewallServiceType makes FirewallServiceType
// nolint
func InterfaceToFirewallServiceType(i interface{}) *FirewallServiceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FirewallServiceType{
		//TODO(nati): Apply default
		Protocol:   common.InterfaceToString(m["protocol"]),
		DSTPorts:   InterfaceToPortType(m["dst_ports"]),
		SRCPorts:   InterfaceToPortType(m["src_ports"]),
		ProtocolID: common.InterfaceToInt64(m["protocol_id"]),
	}
}

// MakeFirewallServiceTypeSlice() makes a slice of FirewallServiceType
// nolint
func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
	return []*FirewallServiceType{}
}

// InterfaceToFirewallServiceTypeSlice() makes a slice of FirewallServiceType
// nolint
func InterfaceToFirewallServiceTypeSlice(i interface{}) []*FirewallServiceType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FirewallServiceType{}
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceType(item))
	}
	return result
}
