package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeDomainLimitsType makes DomainLimitsType
func MakeDomainLimitsType() *DomainLimitsType{
    return &DomainLimitsType{
    //TODO(nati): Apply default
    ProjectLimit: 0,
        VirtualNetworkLimit: 0,
        SecurityGroupLimit: 0,
        
    }
}

// MakeDomainLimitsType makes DomainLimitsType
func InterfaceToDomainLimitsType(i interface{}) *DomainLimitsType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &DomainLimitsType{
    //TODO(nati): Apply default
    ProjectLimit: schema.InterfaceToInt64(m["project_limit"]),
        VirtualNetworkLimit: schema.InterfaceToInt64(m["virtual_network_limit"]),
        SecurityGroupLimit: schema.InterfaceToInt64(m["security_group_limit"]),
        
    }
}

// MakeDomainLimitsTypeSlice() makes a slice of DomainLimitsType
func MakeDomainLimitsTypeSlice() []*DomainLimitsType {
    return []*DomainLimitsType{}
}

// InterfaceToDomainLimitsTypeSlice() makes a slice of DomainLimitsType
func InterfaceToDomainLimitsTypeSlice(i interface{}) []*DomainLimitsType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*DomainLimitsType{}
    for _, item := range list {
        result = append(result, InterfaceToDomainLimitsType(item) )
    }
    return result
}



