package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeVirtualNetworkType makes VirtualNetworkType
func MakeVirtualNetworkType() *VirtualNetworkType{
    return &VirtualNetworkType{
    //TODO(nati): Apply default
    ForwardingMode: "",
        AllowTransit: false,
        NetworkID: 0,
        MirrorDestination: false,
        VxlanNetworkIdentifier: 0,
        RPF: "",
        
    }
}

// MakeVirtualNetworkType makes VirtualNetworkType
func InterfaceToVirtualNetworkType(i interface{}) *VirtualNetworkType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &VirtualNetworkType{
    //TODO(nati): Apply default
    ForwardingMode: schema.InterfaceToString(m["forwarding_mode"]),
        AllowTransit: schema.InterfaceToBool(m["allow_transit"]),
        NetworkID: schema.InterfaceToInt64(m["network_id"]),
        MirrorDestination: schema.InterfaceToBool(m["mirror_destination"]),
        VxlanNetworkIdentifier: schema.InterfaceToInt64(m["vxlan_network_identifier"]),
        RPF: schema.InterfaceToString(m["rpf"]),
        
    }
}

// MakeVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
    return []*VirtualNetworkType{}
}

// InterfaceToVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
func InterfaceToVirtualNetworkTypeSlice(i interface{}) []*VirtualNetworkType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*VirtualNetworkType{}
    for _, item := range list {
        result = append(result, InterfaceToVirtualNetworkType(item) )
    }
    return result
}



