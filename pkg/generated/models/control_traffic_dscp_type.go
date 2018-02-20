package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeControlTrafficDscpType makes ControlTrafficDscpType
func MakeControlTrafficDscpType() *ControlTrafficDscpType{
    return &ControlTrafficDscpType{
    //TODO(nati): Apply default
    Control: 0,
        Analytics: 0,
        DNS: 0,
        
    }
}

// MakeControlTrafficDscpType makes ControlTrafficDscpType
func InterfaceToControlTrafficDscpType(i interface{}) *ControlTrafficDscpType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ControlTrafficDscpType{
    //TODO(nati): Apply default
    Control: schema.InterfaceToInt64(m["control"]),
        Analytics: schema.InterfaceToInt64(m["analytics"]),
        DNS: schema.InterfaceToInt64(m["dns"]),
        
    }
}

// MakeControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
func MakeControlTrafficDscpTypeSlice() []*ControlTrafficDscpType {
    return []*ControlTrafficDscpType{}
}

// InterfaceToControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
func InterfaceToControlTrafficDscpTypeSlice(i interface{}) []*ControlTrafficDscpType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ControlTrafficDscpType{}
    for _, item := range list {
        result = append(result, InterfaceToControlTrafficDscpType(item) )
    }
    return result
}



