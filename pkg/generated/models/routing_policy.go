package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeRoutingPolicy makes RoutingPolicy
func MakeRoutingPolicy() *RoutingPolicy{
    return &RoutingPolicy{
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

// MakeRoutingPolicy makes RoutingPolicy
func InterfaceToRoutingPolicy(i interface{}) *RoutingPolicy{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &RoutingPolicy{
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

// MakeRoutingPolicySlice() makes a slice of RoutingPolicy
func MakeRoutingPolicySlice() []*RoutingPolicy {
    return []*RoutingPolicy{}
}

// InterfaceToRoutingPolicySlice() makes a slice of RoutingPolicy
func InterfaceToRoutingPolicySlice(i interface{}) []*RoutingPolicy {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*RoutingPolicy{}
    for _, item := range list {
        result = append(result, InterfaceToRoutingPolicy(item) )
    }
    return result
}



