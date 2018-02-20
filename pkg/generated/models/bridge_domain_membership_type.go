package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeBridgeDomainMembershipType makes BridgeDomainMembershipType
func MakeBridgeDomainMembershipType() *BridgeDomainMembershipType{
    return &BridgeDomainMembershipType{
    //TODO(nati): Apply default
    VlanTag: 0,
        
    }
}

// MakeBridgeDomainMembershipType makes BridgeDomainMembershipType
func InterfaceToBridgeDomainMembershipType(i interface{}) *BridgeDomainMembershipType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &BridgeDomainMembershipType{
    //TODO(nati): Apply default
    VlanTag: schema.InterfaceToInt64(m["vlan_tag"]),
        
    }
}

// MakeBridgeDomainMembershipTypeSlice() makes a slice of BridgeDomainMembershipType
func MakeBridgeDomainMembershipTypeSlice() []*BridgeDomainMembershipType {
    return []*BridgeDomainMembershipType{}
}

// InterfaceToBridgeDomainMembershipTypeSlice() makes a slice of BridgeDomainMembershipType
func InterfaceToBridgeDomainMembershipTypeSlice(i interface{}) []*BridgeDomainMembershipType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*BridgeDomainMembershipType{}
    for _, item := range list {
        result = append(result, InterfaceToBridgeDomainMembershipType(item) )
    }
    return result
}



