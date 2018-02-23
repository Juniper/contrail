package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakePolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleType() *PolicyBasedForwardingRuleType{
    return &PolicyBasedForwardingRuleType{
    //TODO(nati): Apply default
    DSTMac: "",
        Protocol: "",
        Ipv6ServiceChainAddress: "",
        Direction: "",
        MPLSLabel: 0,
        VlanTag: 0,
        SRCMac: "",
        ServiceChainAddress: "",
        
    }
}

// MakePolicyBasedForwardingRuleType makes PolicyBasedForwardingRuleType
func InterfaceToPolicyBasedForwardingRuleType(i interface{}) *PolicyBasedForwardingRuleType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &PolicyBasedForwardingRuleType{
    //TODO(nati): Apply default
    DSTMac: schema.InterfaceToString(m["dst_mac"]),
        Protocol: schema.InterfaceToString(m["protocol"]),
        Ipv6ServiceChainAddress: schema.InterfaceToString(m["ipv6_service_chain_address"]),
        Direction: schema.InterfaceToString(m["direction"]),
        MPLSLabel: schema.InterfaceToInt64(m["mpls_label"]),
        VlanTag: schema.InterfaceToInt64(m["vlan_tag"]),
        SRCMac: schema.InterfaceToString(m["src_mac"]),
        ServiceChainAddress: schema.InterfaceToString(m["service_chain_address"]),
        
    }
}

// MakePolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
func MakePolicyBasedForwardingRuleTypeSlice() []*PolicyBasedForwardingRuleType {
    return []*PolicyBasedForwardingRuleType{}
}

// InterfaceToPolicyBasedForwardingRuleTypeSlice() makes a slice of PolicyBasedForwardingRuleType
func InterfaceToPolicyBasedForwardingRuleTypeSlice(i interface{}) []*PolicyBasedForwardingRuleType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*PolicyBasedForwardingRuleType{}
    for _, item := range list {
        result = append(result, InterfaceToPolicyBasedForwardingRuleType(item) )
    }
    return result
}



