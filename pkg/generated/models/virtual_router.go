package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeVirtualRouter makes VirtualRouter
func MakeVirtualRouter() *VirtualRouter{
    return &VirtualRouter{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        VirtualRouterDPDKEnabled: false,
        VirtualRouterType: "",
        VirtualRouterIPAddress: "",
        
    }
}

// MakeVirtualRouter makes VirtualRouter
func InterfaceToVirtualRouter(i interface{}) *VirtualRouter{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &VirtualRouter{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        VirtualRouterDPDKEnabled: schema.InterfaceToBool(m["virtual_router_dpdk_enabled"]),
        VirtualRouterType: schema.InterfaceToString(m["virtual_router_type"]),
        VirtualRouterIPAddress: schema.InterfaceToString(m["virtual_router_ip_address"]),
        
    }
}

// MakeVirtualRouterSlice() makes a slice of VirtualRouter
func MakeVirtualRouterSlice() []*VirtualRouter {
    return []*VirtualRouter{}
}

// InterfaceToVirtualRouterSlice() makes a slice of VirtualRouter
func InterfaceToVirtualRouterSlice(i interface{}) []*VirtualRouter {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*VirtualRouter{}
    for _, item := range list {
        result = append(result, InterfaceToVirtualRouter(item) )
    }
    return result
}



