package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLogicalInterface makes LogicalInterface
func MakeLogicalInterface() *LogicalInterface{
    return &LogicalInterface{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        LogicalInterfaceVlanTag: 0,
        LogicalInterfaceType: "",
        
    }
}

// MakeLogicalInterface makes LogicalInterface
func InterfaceToLogicalInterface(i interface{}) *LogicalInterface{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &LogicalInterface{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        LogicalInterfaceVlanTag: schema.InterfaceToInt64(m["logical_interface_vlan_tag"]),
        LogicalInterfaceType: schema.InterfaceToString(m["logical_interface_type"]),
        
    }
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
    return []*LogicalInterface{}
}

// InterfaceToLogicalInterfaceSlice() makes a slice of LogicalInterface
func InterfaceToLogicalInterfaceSlice(i interface{}) []*LogicalInterface {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*LogicalInterface{}
    for _, item := range list {
        result = append(result, InterfaceToLogicalInterface(item) )
    }
    return result
}



