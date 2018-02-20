package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeVirtualMachine makes VirtualMachine
func MakeVirtualMachine() *VirtualMachine{
    return &VirtualMachine{
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

// MakeVirtualMachine makes VirtualMachine
func InterfaceToVirtualMachine(i interface{}) *VirtualMachine{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &VirtualMachine{
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

// MakeVirtualMachineSlice() makes a slice of VirtualMachine
func MakeVirtualMachineSlice() []*VirtualMachine {
    return []*VirtualMachine{}
}

// InterfaceToVirtualMachineSlice() makes a slice of VirtualMachine
func InterfaceToVirtualMachineSlice(i interface{}) []*VirtualMachine {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*VirtualMachine{}
    for _, item := range list {
        result = append(result, InterfaceToVirtualMachine(item) )
    }
    return result
}



