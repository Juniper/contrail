package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeRouteTarget makes RouteTarget
func MakeRouteTarget() *RouteTarget{
    return &RouteTarget{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
    }
}

// MakeRouteTarget makes RouteTarget
func InterfaceToRouteTarget(i interface{}) *RouteTarget{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &RouteTarget{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        
    }
}

// MakeRouteTargetSlice() makes a slice of RouteTarget
func MakeRouteTargetSlice() []*RouteTarget {
    return []*RouteTarget{}
}

// InterfaceToRouteTargetSlice() makes a slice of RouteTarget
func InterfaceToRouteTargetSlice(i interface{}) []*RouteTarget {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*RouteTarget{}
    for _, item := range list {
        result = append(result, InterfaceToRouteTarget(item) )
    }
    return result
}



