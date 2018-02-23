package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeInstanceIP makes InstanceIP
func MakeInstanceIP() *InstanceIP{
    return &InstanceIP{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceHealthCheckIP: false,
        SecondaryIPTrackingIP: MakeSubnetType(),
        InstanceIPAddress: "",
        InstanceIPMode: "",
        SubnetUUID: "",
        InstanceIPFamily: "",
        ServiceInstanceIP: false,
        InstanceIPLocalIP: false,
        InstanceIPSecondary: false,
        
    }
}

// MakeInstanceIP makes InstanceIP
func InterfaceToInstanceIP(i interface{}) *InstanceIP{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &InstanceIP{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        ServiceHealthCheckIP: schema.InterfaceToBool(m["service_health_check_ip"]),
        SecondaryIPTrackingIP: InterfaceToSubnetType(m["secondary_ip_tracking_ip"]),
        InstanceIPAddress: schema.InterfaceToString(m["instance_ip_address"]),
        InstanceIPMode: schema.InterfaceToString(m["instance_ip_mode"]),
        SubnetUUID: schema.InterfaceToString(m["subnet_uuid"]),
        InstanceIPFamily: schema.InterfaceToString(m["instance_ip_family"]),
        ServiceInstanceIP: schema.InterfaceToBool(m["service_instance_ip"]),
        InstanceIPLocalIP: schema.InterfaceToBool(m["instance_ip_local_ip"]),
        InstanceIPSecondary: schema.InterfaceToBool(m["instance_ip_secondary"]),
        
    }
}

// MakeInstanceIPSlice() makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
    return []*InstanceIP{}
}

// InterfaceToInstanceIPSlice() makes a slice of InstanceIP
func InterfaceToInstanceIPSlice(i interface{}) []*InstanceIP {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*InstanceIP{}
    for _, item := range list {
        result = append(result, InterfaceToInstanceIP(item) )
    }
    return result
}



