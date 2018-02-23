package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeServiceHealthCheck makes ServiceHealthCheck
func MakeServiceHealthCheck() *ServiceHealthCheck{
    return &ServiceHealthCheck{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceHealthCheckProperties: MakeServiceHealthCheckType(),
        
    }
}

// MakeServiceHealthCheck makes ServiceHealthCheck
func InterfaceToServiceHealthCheck(i interface{}) *ServiceHealthCheck{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceHealthCheck{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        ServiceHealthCheckProperties: InterfaceToServiceHealthCheckType(m["service_health_check_properties"]),
        
    }
}

// MakeServiceHealthCheckSlice() makes a slice of ServiceHealthCheck
func MakeServiceHealthCheckSlice() []*ServiceHealthCheck {
    return []*ServiceHealthCheck{}
}

// InterfaceToServiceHealthCheckSlice() makes a slice of ServiceHealthCheck
func InterfaceToServiceHealthCheckSlice(i interface{}) []*ServiceHealthCheck {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceHealthCheck{}
    for _, item := range list {
        result = append(result, InterfaceToServiceHealthCheck(item) )
    }
    return result
}



