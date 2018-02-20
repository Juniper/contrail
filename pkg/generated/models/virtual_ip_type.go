package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeVirtualIpType makes VirtualIpType
func MakeVirtualIpType() *VirtualIpType{
    return &VirtualIpType{
    //TODO(nati): Apply default
    Status: "",
        StatusDescription: "",
        Protocol: "",
        SubnetID: "",
        PersistenceCookieName: "",
        ConnectionLimit: 0,
        PersistenceType: "",
        AdminState: false,
        Address: "",
        ProtocolPort: 0,
        
    }
}

// MakeVirtualIpType makes VirtualIpType
func InterfaceToVirtualIpType(i interface{}) *VirtualIpType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &VirtualIpType{
    //TODO(nati): Apply default
    Status: schema.InterfaceToString(m["status"]),
        StatusDescription: schema.InterfaceToString(m["status_description"]),
        Protocol: schema.InterfaceToString(m["protocol"]),
        SubnetID: schema.InterfaceToString(m["subnet_id"]),
        PersistenceCookieName: schema.InterfaceToString(m["persistence_cookie_name"]),
        ConnectionLimit: schema.InterfaceToInt64(m["connection_limit"]),
        PersistenceType: schema.InterfaceToString(m["persistence_type"]),
        AdminState: schema.InterfaceToBool(m["admin_state"]),
        Address: schema.InterfaceToString(m["address"]),
        ProtocolPort: schema.InterfaceToInt64(m["protocol_port"]),
        
    }
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
    return []*VirtualIpType{}
}

// InterfaceToVirtualIpTypeSlice() makes a slice of VirtualIpType
func InterfaceToVirtualIpTypeSlice(i interface{}) []*VirtualIpType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*VirtualIpType{}
    for _, item := range list {
        result = append(result, InterfaceToVirtualIpType(item) )
    }
    return result
}



