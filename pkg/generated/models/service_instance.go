package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeServiceInstance makes ServiceInstance
func MakeServiceInstance() *ServiceInstance{
    return &ServiceInstance{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ServiceInstanceBindings: MakeKeyValuePairs(),
        ServiceInstanceProperties: MakeServiceInstanceType(),
        
    }
}

// MakeServiceInstance makes ServiceInstance
func InterfaceToServiceInstance(i interface{}) *ServiceInstance{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceInstance{
    //TODO(nati): Apply default
    UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        ServiceInstanceBindings: InterfaceToKeyValuePairs(m["service_instance_bindings"]),
        ServiceInstanceProperties: InterfaceToServiceInstanceType(m["service_instance_properties"]),
        
    }
}

// MakeServiceInstanceSlice() makes a slice of ServiceInstance
func MakeServiceInstanceSlice() []*ServiceInstance {
    return []*ServiceInstance{}
}

// InterfaceToServiceInstanceSlice() makes a slice of ServiceInstance
func InterfaceToServiceInstanceSlice(i interface{}) []*ServiceInstance {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceInstance{}
    for _, item := range list {
        result = append(result, InterfaceToServiceInstance(item) )
    }
    return result
}



