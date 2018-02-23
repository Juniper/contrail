package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeUser makes User
func MakeUser() *User{
    return &User{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Password: "",
        
    }
}

// MakeUser makes User
func InterfaceToUser(i interface{}) *User{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &User{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        Password: schema.InterfaceToString(m["password"]),
        
    }
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
    return []*User{}
}

// InterfaceToUserSlice() makes a slice of User
func InterfaceToUserSlice(i interface{}) []*User {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*User{}
    for _, item := range list {
        result = append(result, InterfaceToUser(item) )
    }
    return result
}



