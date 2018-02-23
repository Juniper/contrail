package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLoadbalancerMemberType makes LoadbalancerMemberType
func MakeLoadbalancerMemberType() *LoadbalancerMemberType{
    return &LoadbalancerMemberType{
    //TODO(nati): Apply default
    Status: "",
        StatusDescription: "",
        Weight: 0,
        AdminState: false,
        Address: "",
        ProtocolPort: 0,
        
    }
}

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
func InterfaceToLoadbalancerMemberType(i interface{}) *LoadbalancerMemberType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &LoadbalancerMemberType{
    //TODO(nati): Apply default
    Status: schema.InterfaceToString(m["status"]),
        StatusDescription: schema.InterfaceToString(m["status_description"]),
        Weight: schema.InterfaceToInt64(m["weight"]),
        AdminState: schema.InterfaceToBool(m["admin_state"]),
        Address: schema.InterfaceToString(m["address"]),
        ProtocolPort: schema.InterfaceToInt64(m["protocol_port"]),
        
    }
}

// MakeLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
    return []*LoadbalancerMemberType{}
}

// InterfaceToLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
func InterfaceToLoadbalancerMemberTypeSlice(i interface{}) []*LoadbalancerMemberType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*LoadbalancerMemberType{}
    for _, item := range list {
        result = append(result, InterfaceToLoadbalancerMemberType(item) )
    }
    return result
}



