package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType
// nolint
func MakePolicyBasedForwardingRuleType() *PolicyBasedForwardingRuleType {
	return &PolicyBasedForwardingRuleType{
		//TODO(nati): Apply default
		DSTMac:                  "",
		Protocol:                "",
		Ipv6ServiceChainAddress: "",
		Direction:               "",
		MPLSLabel:               0,
		VlanTag:                 0,
		SRCMac:                  "",
		ServiceChainAddress:     "",
	}
}

// MakePolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType
// nolint
func InterfaceToPolicyBasedForwardingRuleType(i interface{}) *PolicyBasedForwardingRuleType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PolicyBasedForwardingRuleType{
		//TODO(nati): Apply default
		DSTMac:                  common.InterfaceToString(m["dst_mac"]),
		Protocol:                common.InterfaceToString(m["protocol"]),
		Ipv6ServiceChainAddress: common.InterfaceToString(m["ipv6_service_chain_address"]),
		Direction:               common.InterfaceToString(m["direction"]),
		MPLSLabel:               common.InterfaceToInt64(m["mpls_label"]),
		VlanTag:                 common.InterfaceToInt64(m["vlan_tag"]),
		SRCMac:                  common.InterfaceToString(m["src_mac"]),
		ServiceChainAddress:     common.InterfaceToString(m["service_chain_address"]),
	}
}

// MakePolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
// nolint
func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
	return []*PolicyBasedForwardingRuleType{}
}

// InterfaceToPolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
// nolint
func InterfaceToPolicyBasedForwardingRuleTypeSlice(i interface{}) []*PolicyBasedForwardingRuleType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PolicyBasedForwardingRuleType{}
	for _, item := range list {
		result = append(result, InterfaceToPolicyBasedForwardingRuleType(item))
	}
	return result
}
