package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeE2ServiceProvider makes E2ServiceProvider
func MakeE2ServiceProvider() *E2ServiceProvider{
    return &E2ServiceProvider{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        E2ServiceProviderPromiscuous: false,
        
    }
}

// MakeE2ServiceProvider makes E2ServiceProvider
func InterfaceToE2ServiceProvider(i interface{}) *E2ServiceProvider{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &E2ServiceProvider{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        E2ServiceProviderPromiscuous: schema.InterfaceToBool(m["e2_service_provider_promiscuous"]),
        
    }
}

// MakeE2ServiceProviderSlice() makes a slice of E2ServiceProvider
func MakeE2ServiceProviderSlice() []*E2ServiceProvider {
    return []*E2ServiceProvider{}
}

// InterfaceToE2ServiceProviderSlice() makes a slice of E2ServiceProvider
func InterfaceToE2ServiceProviderSlice(i interface{}) []*E2ServiceProvider {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*E2ServiceProvider{}
    for _, item := range list {
        result = append(result, InterfaceToE2ServiceProvider(item) )
    }
    return result
}



