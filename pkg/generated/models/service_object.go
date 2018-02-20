package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeServiceObject makes ServiceObject
func MakeServiceObject() *ServiceObject{
    return &ServiceObject{
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

// MakeServiceObject makes ServiceObject
func InterfaceToServiceObject(i interface{}) *ServiceObject{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceObject{
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

// MakeServiceObjectSlice() makes a slice of ServiceObject
func MakeServiceObjectSlice() []*ServiceObject {
    return []*ServiceObject{}
}

// InterfaceToServiceObjectSlice() makes a slice of ServiceObject
func InterfaceToServiceObjectSlice(i interface{}) []*ServiceObject {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceObject{}
    for _, item := range list {
        result = append(result, InterfaceToServiceObject(item) )
    }
    return result
}



