package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLoadbalancerPool makes LoadbalancerPool
func MakeLoadbalancerPool() *LoadbalancerPool{
    return &LoadbalancerPool{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        LoadbalancerPoolProperties: MakeLoadbalancerPoolType(),
        LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
        LoadbalancerPoolProvider: "",
        
    }
}

// MakeLoadbalancerPool makes LoadbalancerPool
func InterfaceToLoadbalancerPool(i interface{}) *LoadbalancerPool{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &LoadbalancerPool{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        LoadbalancerPoolProperties: InterfaceToLoadbalancerPoolType(m["loadbalancer_pool_properties"]),
        LoadbalancerPoolCustomAttributes: InterfaceToKeyValuePairs(m["loadbalancer_pool_custom_attributes"]),
        LoadbalancerPoolProvider: schema.InterfaceToString(m["loadbalancer_pool_provider"]),
        
    }
}

// MakeLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
    return []*LoadbalancerPool{}
}

// InterfaceToLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
func InterfaceToLoadbalancerPoolSlice(i interface{}) []*LoadbalancerPool {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*LoadbalancerPool{}
    for _, item := range list {
        result = append(result, InterfaceToLoadbalancerPool(item) )
    }
    return result
}



