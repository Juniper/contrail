package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeRouteTable makes RouteTable
func MakeRouteTable() *RouteTable{
    return &RouteTable{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Routes: MakeRouteTableType(),
        
    }
}

// MakeRouteTable makes RouteTable
func InterfaceToRouteTable(i interface{}) *RouteTable{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &RouteTable{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        Routes: InterfaceToRouteTableType(m["routes"]),
        
    }
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
    return []*RouteTable{}
}

// InterfaceToRouteTableSlice() makes a slice of RouteTable
func InterfaceToRouteTableSlice(i interface{}) []*RouteTable {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*RouteTable{}
    for _, item := range list {
        result = append(result, InterfaceToRouteTable(item) )
    }
    return result
}



