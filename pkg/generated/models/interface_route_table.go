package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeInterfaceRouteTable makes InterfaceRouteTable
func MakeInterfaceRouteTable() *InterfaceRouteTable{
    return &InterfaceRouteTable{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        InterfaceRouteTableRoutes: MakeRouteTableType(),
        
    }
}

// MakeInterfaceRouteTable makes InterfaceRouteTable
func InterfaceToInterfaceRouteTable(i interface{}) *InterfaceRouteTable{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &InterfaceRouteTable{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        InterfaceRouteTableRoutes: InterfaceToRouteTableType(m["interface_route_table_routes"]),
        
    }
}

// MakeInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
func MakeInterfaceRouteTableSlice() []*InterfaceRouteTable {
    return []*InterfaceRouteTable{}
}

// InterfaceToInterfaceRouteTableSlice() makes a slice of InterfaceRouteTable
func InterfaceToInterfaceRouteTableSlice(i interface{}) []*InterfaceRouteTable {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*InterfaceRouteTable{}
    for _, item := range list {
        result = append(result, InterfaceToInterfaceRouteTable(item) )
    }
    return result
}



