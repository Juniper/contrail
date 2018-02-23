package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeBaremetalPort makes BaremetalPort
func MakeBaremetalPort() *BaremetalPort{
    return &BaremetalPort{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        MacAddress: "",
        CreatedAt: "",
        UpdatedAt: "",
        Node: "",
        PxeEnabled: false,
        LocalLinkConnection: MakeLocalLinkConnection(),
        
    }
}

// MakeBaremetalPort makes BaremetalPort
func InterfaceToBaremetalPort(i interface{}) *BaremetalPort{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &BaremetalPort{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        MacAddress: schema.InterfaceToString(m["mac_address"]),
        CreatedAt: schema.InterfaceToString(m["created_at"]),
        UpdatedAt: schema.InterfaceToString(m["updated_at"]),
        Node: schema.InterfaceToString(m["node"]),
        PxeEnabled: schema.InterfaceToBool(m["pxe_enabled"]),
        LocalLinkConnection: InterfaceToLocalLinkConnection(m["local_link_connection"]),
        
    }
}

// MakeBaremetalPortSlice() makes a slice of BaremetalPort
func MakeBaremetalPortSlice() []*BaremetalPort {
    return []*BaremetalPort{}
}

// InterfaceToBaremetalPortSlice() makes a slice of BaremetalPort
func InterfaceToBaremetalPortSlice(i interface{}) []*BaremetalPort {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*BaremetalPort{}
    for _, item := range list {
        result = append(result, InterfaceToBaremetalPort(item) )
    }
    return result
}



