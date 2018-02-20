package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeNetworkDeviceConfig makes NetworkDeviceConfig
func MakeNetworkDeviceConfig() *NetworkDeviceConfig{
    return &NetworkDeviceConfig{
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

// MakeNetworkDeviceConfig makes NetworkDeviceConfig
func InterfaceToNetworkDeviceConfig(i interface{}) *NetworkDeviceConfig{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &NetworkDeviceConfig{
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

// MakeNetworkDeviceConfigSlice() makes a slice of NetworkDeviceConfig
func MakeNetworkDeviceConfigSlice() []*NetworkDeviceConfig {
    return []*NetworkDeviceConfig{}
}

// InterfaceToNetworkDeviceConfigSlice() makes a slice of NetworkDeviceConfig
func InterfaceToNetworkDeviceConfigSlice(i interface{}) []*NetworkDeviceConfig {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*NetworkDeviceConfig{}
    for _, item := range list {
        result = append(result, InterfaceToNetworkDeviceConfig(item) )
    }
    return result
}



