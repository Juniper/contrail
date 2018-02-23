package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeMacAddressesType makes MacAddressesType
func MakeMacAddressesType() *MacAddressesType{
    return &MacAddressesType{
    //TODO(nati): Apply default
    MacAddress: []string{},
        
    }
}

// MakeMacAddressesType makes MacAddressesType
func InterfaceToMacAddressesType(i interface{}) *MacAddressesType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &MacAddressesType{
    //TODO(nati): Apply default
    MacAddress: schema.InterfaceToStringList(m["mac_address"]),
        
    }
}

// MakeMacAddressesTypeSlice() makes a slice of MacAddressesType
func MakeMacAddressesTypeSlice() []*MacAddressesType {
    return []*MacAddressesType{}
}

// InterfaceToMacAddressesTypeSlice() makes a slice of MacAddressesType
func InterfaceToMacAddressesTypeSlice(i interface{}) []*MacAddressesType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*MacAddressesType{}
    for _, item := range list {
        result = append(result, InterfaceToMacAddressesType(item) )
    }
    return result
}



