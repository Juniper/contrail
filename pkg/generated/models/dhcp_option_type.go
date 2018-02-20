package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeDhcpOptionType makes DhcpOptionType
func MakeDhcpOptionType() *DhcpOptionType{
    return &DhcpOptionType{
    //TODO(nati): Apply default
    DHCPOptionValue: "",
        DHCPOptionValueBytes: "",
        DHCPOptionName: "",
        
    }
}

// MakeDhcpOptionType makes DhcpOptionType
func InterfaceToDhcpOptionType(i interface{}) *DhcpOptionType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &DhcpOptionType{
    //TODO(nati): Apply default
    DHCPOptionValue: schema.InterfaceToString(m["dhcp_option_value"]),
        DHCPOptionValueBytes: schema.InterfaceToString(m["dhcp_option_value_bytes"]),
        DHCPOptionName: schema.InterfaceToString(m["dhcp_option_name"]),
        
    }
}

// MakeDhcpOptionTypeSlice() makes a slice of DhcpOptionType
func MakeDhcpOptionTypeSlice() []*DhcpOptionType {
    return []*DhcpOptionType{}
}

// InterfaceToDhcpOptionTypeSlice() makes a slice of DhcpOptionType
func InterfaceToDhcpOptionTypeSlice(i interface{}) []*DhcpOptionType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*DhcpOptionType{}
    for _, item := range list {
        result = append(result, InterfaceToDhcpOptionType(item) )
    }
    return result
}



