package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAllowedAddressPair makes AllowedAddressPair
func MakeAllowedAddressPair() *AllowedAddressPair{
    return &AllowedAddressPair{
    //TODO(nati): Apply default
    IP: MakeSubnetType(),
        Mac: "",
        AddressMode: "",
        
    }
}

// MakeAllowedAddressPair makes AllowedAddressPair
func InterfaceToAllowedAddressPair(i interface{}) *AllowedAddressPair{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AllowedAddressPair{
    //TODO(nati): Apply default
    IP: InterfaceToSubnetType(m["ip"]),
        Mac: schema.InterfaceToString(m["mac"]),
        AddressMode: schema.InterfaceToString(m["address_mode"]),
        
    }
}

// MakeAllowedAddressPairSlice() makes a slice of AllowedAddressPair
func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
    return []*AllowedAddressPair{}
}

// InterfaceToAllowedAddressPairSlice() makes a slice of AllowedAddressPair
func InterfaceToAllowedAddressPairSlice(i interface{}) []*AllowedAddressPair {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AllowedAddressPair{}
    for _, item := range list {
        result = append(result, InterfaceToAllowedAddressPair(item) )
    }
    return result
}



