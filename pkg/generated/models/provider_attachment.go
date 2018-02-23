package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeProviderAttachment makes ProviderAttachment
func MakeProviderAttachment() *ProviderAttachment{
    return &ProviderAttachment{
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

// MakeProviderAttachment makes ProviderAttachment
func InterfaceToProviderAttachment(i interface{}) *ProviderAttachment{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ProviderAttachment{
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

// MakeProviderAttachmentSlice() makes a slice of ProviderAttachment
func MakeProviderAttachmentSlice() []*ProviderAttachment {
    return []*ProviderAttachment{}
}

// InterfaceToProviderAttachmentSlice() makes a slice of ProviderAttachment
func InterfaceToProviderAttachmentSlice(i interface{}) []*ProviderAttachment {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ProviderAttachment{}
    for _, item := range list {
        result = append(result, InterfaceToProviderAttachment(item) )
    }
    return result
}



