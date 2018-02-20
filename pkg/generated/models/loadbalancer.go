package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLoadbalancer makes Loadbalancer
func MakeLoadbalancer() *Loadbalancer{
    return &Loadbalancer{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        LoadbalancerProperties: MakeLoadbalancerType(),
        LoadbalancerProvider: "",
        
    }
}

// MakeLoadbalancer makes Loadbalancer
func InterfaceToLoadbalancer(i interface{}) *Loadbalancer{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &Loadbalancer{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        LoadbalancerProperties: InterfaceToLoadbalancerType(m["loadbalancer_properties"]),
        LoadbalancerProvider: schema.InterfaceToString(m["loadbalancer_provider"]),
        
    }
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
    return []*Loadbalancer{}
}

// InterfaceToLoadbalancerSlice() makes a slice of Loadbalancer
func InterfaceToLoadbalancerSlice(i interface{}) []*Loadbalancer {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*Loadbalancer{}
    for _, item := range list {
        result = append(result, InterfaceToLoadbalancer(item) )
    }
    return result
}



