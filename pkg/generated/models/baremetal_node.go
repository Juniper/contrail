package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeBaremetalNode makes BaremetalNode
func MakeBaremetalNode() *BaremetalNode{
    return &BaremetalNode{
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
        DriverInfo: MakeDriverInfo(),
        BMProperties: MakeBaremetalProperties(),
        InstanceUUID: "",
        InstanceInfo: MakeInstanceInfo(),
        Maintenance: false,
        MaintenanceReason: "",
        PowerState: "",
        TargetPowerState: "",
        ProvisionState: "",
        TargetProvisionState: "",
        ConsoleEnabled: false,
        CreatedAt: "",
        UpdatedAt: "",
        LastError: "",
        
    }
}

// MakeBaremetalNode makes BaremetalNode
func InterfaceToBaremetalNode(i interface{}) *BaremetalNode{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &BaremetalNode{
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
        DriverInfo: InterfaceToDriverInfo(m["driver_info"]),
        BMProperties: InterfaceToBaremetalProperties(m["bm_properties"]),
        InstanceUUID: schema.InterfaceToString(m["instance_uuid"]),
        InstanceInfo: InterfaceToInstanceInfo(m["instance_info"]),
        Maintenance: schema.InterfaceToBool(m["maintenance"]),
        MaintenanceReason: schema.InterfaceToString(m["maintenance_reason"]),
        PowerState: schema.InterfaceToString(m["power_state"]),
        TargetPowerState: schema.InterfaceToString(m["target_power_state"]),
        ProvisionState: schema.InterfaceToString(m["provision_state"]),
        TargetProvisionState: schema.InterfaceToString(m["target_provision_state"]),
        ConsoleEnabled: schema.InterfaceToBool(m["console_enabled"]),
        CreatedAt: schema.InterfaceToString(m["created_at"]),
        UpdatedAt: schema.InterfaceToString(m["updated_at"]),
        LastError: schema.InterfaceToString(m["last_error"]),
        
    }
}

// MakeBaremetalNodeSlice() makes a slice of BaremetalNode
func MakeBaremetalNodeSlice() []*BaremetalNode {
    return []*BaremetalNode{}
}

// InterfaceToBaremetalNodeSlice() makes a slice of BaremetalNode
func InterfaceToBaremetalNodeSlice(i interface{}) []*BaremetalNode {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*BaremetalNode{}
    for _, item := range list {
        result = append(result, InterfaceToBaremetalNode(item) )
    }
    return result
}



