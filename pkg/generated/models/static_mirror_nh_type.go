package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeStaticMirrorNhType makes StaticMirrorNhType
func MakeStaticMirrorNhType() *StaticMirrorNhType{
    return &StaticMirrorNhType{
    //TODO(nati): Apply default
    VtepDSTIPAddress: "",
        VtepDSTMacAddress: "",
        Vni: 0,
        
    }
}

// MakeStaticMirrorNhType makes StaticMirrorNhType
func InterfaceToStaticMirrorNhType(i interface{}) *StaticMirrorNhType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &StaticMirrorNhType{
    //TODO(nati): Apply default
    VtepDSTIPAddress: schema.InterfaceToString(m["vtep_dst_ip_address"]),
        VtepDSTMacAddress: schema.InterfaceToString(m["vtep_dst_mac_address"]),
        Vni: schema.InterfaceToInt64(m["vni"]),
        
    }
}

// MakeStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
func MakeStaticMirrorNhTypeSlice() []*StaticMirrorNhType {
    return []*StaticMirrorNhType{}
}

// InterfaceToStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
func InterfaceToStaticMirrorNhTypeSlice(i interface{}) []*StaticMirrorNhType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*StaticMirrorNhType{}
    for _, item := range list {
        result = append(result, InterfaceToStaticMirrorNhType(item) )
    }
    return result
}



