package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeTag makes Tag
func MakeTag() *Tag{
    return &Tag{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        TagTypeName: "",
        TagID: "",
        TagValue: "",
        
    }
}

// MakeTag makes Tag
func InterfaceToTag(i interface{}) *Tag{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &Tag{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        TagTypeName: schema.InterfaceToString(m["tag_type_name"]),
        TagID: schema.InterfaceToString(m["tag_id"]),
        TagValue: schema.InterfaceToString(m["tag_value"]),
        
    }
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
    return []*Tag{}
}

// InterfaceToTagSlice() makes a slice of Tag
func InterfaceToTagSlice(i interface{}) []*Tag {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*Tag{}
    for _, item := range list {
        result = append(result, InterfaceToTag(item) )
    }
    return result
}



