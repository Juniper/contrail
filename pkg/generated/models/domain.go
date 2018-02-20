package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeDomain makes Domain
func MakeDomain() *Domain{
    return &Domain{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        DomainLimits: MakeDomainLimitsType(),
        
    }
}

// MakeDomain makes Domain
func InterfaceToDomain(i interface{}) *Domain{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &Domain{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        DomainLimits: InterfaceToDomainLimitsType(m["domain_limits"]),
        
    }
}

// MakeDomainSlice() makes a slice of Domain
func MakeDomainSlice() []*Domain {
    return []*Domain{}
}

// InterfaceToDomainSlice() makes a slice of Domain
func InterfaceToDomainSlice(i interface{}) []*Domain {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*Domain{}
    for _, item := range list {
        result = append(result, InterfaceToDomain(item) )
    }
    return result
}



