package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakePortMap makes PortMap
func MakePortMap() *PortMap{
    return &PortMap{
    //TODO(nati): Apply default
    SRCPort: 0,
        Protocol: "",
        DSTPort: 0,
        
    }
}

// MakePortMap makes PortMap
func InterfaceToPortMap(i interface{}) *PortMap{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &PortMap{
    //TODO(nati): Apply default
    SRCPort: schema.InterfaceToInt64(m["src_port"]),
        Protocol: schema.InterfaceToString(m["protocol"]),
        DSTPort: schema.InterfaceToInt64(m["dst_port"]),
        
    }
}

// MakePortMapSlice() makes a slice of PortMap
func MakePortMapSlice() []*PortMap {
    return []*PortMap{}
}

// InterfaceToPortMapSlice() makes a slice of PortMap
func InterfaceToPortMapSlice(i interface{}) []*PortMap {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*PortMap{}
    for _, item := range list {
        result = append(result, InterfaceToPortMap(item) )
    }
    return result
}



