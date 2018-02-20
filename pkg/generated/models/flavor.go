package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeFlavor makes Flavor
func MakeFlavor() *Flavor{
    return &Flavor{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Name: "",
        Disk: 0,
        Vcpus: 0,
        RAM: 0,
        ID: "",
        Property: "",
        RXTXFactor: 0,
        Swap: 0,
        IsPublic: false,
        Ephemeral: 0,
        Links: MakeOpenStackLink(),
        
    }
}

// MakeFlavor makes Flavor
func InterfaceToFlavor(i interface{}) *Flavor{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &Flavor{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        Name: schema.InterfaceToString(m["name"]),
        Disk: schema.InterfaceToInt64(m["disk"]),
        Vcpus: schema.InterfaceToInt64(m["vcpus"]),
        RAM: schema.InterfaceToInt64(m["ram"]),
        ID: schema.InterfaceToString(m["id"]),
        Property: schema.InterfaceToString(m["property"]),
        RXTXFactor: schema.InterfaceToInt64(m["rxtx_factor"]),
        Swap: schema.InterfaceToInt64(m["swap"]),
        IsPublic: schema.InterfaceToBool(m["is_public"]),
        Ephemeral: schema.InterfaceToInt64(m["ephemeral"]),
        Links: InterfaceToOpenStackLink(m["links"]),
        
    }
}

// MakeFlavorSlice() makes a slice of Flavor
func MakeFlavorSlice() []*Flavor {
    return []*Flavor{}
}

// InterfaceToFlavorSlice() makes a slice of Flavor
func InterfaceToFlavorSlice(i interface{}) []*Flavor {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*Flavor{}
    for _, item := range list {
        result = append(result, InterfaceToFlavor(item) )
    }
    return result
}



