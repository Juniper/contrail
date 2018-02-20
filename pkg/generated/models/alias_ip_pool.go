package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAliasIPPool makes AliasIPPool
func MakeAliasIPPool() *AliasIPPool{
    return &AliasIPPool{
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

// MakeAliasIPPool makes AliasIPPool
func InterfaceToAliasIPPool(i interface{}) *AliasIPPool{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AliasIPPool{
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

// MakeAliasIPPoolSlice() makes a slice of AliasIPPool
func MakeAliasIPPoolSlice() []*AliasIPPool {
    return []*AliasIPPool{}
}

// InterfaceToAliasIPPoolSlice() makes a slice of AliasIPPool
func InterfaceToAliasIPPoolSlice(i interface{}) []*AliasIPPool {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AliasIPPool{}
    for _, item := range list {
        result = append(result, InterfaceToAliasIPPool(item) )
    }
    return result
}



