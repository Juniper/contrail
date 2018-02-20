package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeInstanceInfo makes InstanceInfo
func MakeInstanceInfo() *InstanceInfo{
    return &InstanceInfo{
    //TODO(nati): Apply default
    DisplayName: "",
        ImageSource: "",
        LocalGB: "",
        MemoryMB: "",
        NovaHostID: "",
        RootGB: "",
        SwapMB: "",
        Vcpus: "",
        Capabilities: "",
        
    }
}

// MakeInstanceInfo makes InstanceInfo
func InterfaceToInstanceInfo(i interface{}) *InstanceInfo{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &InstanceInfo{
    //TODO(nati): Apply default
    DisplayName: schema.InterfaceToString(m["display_name"]),
        ImageSource: schema.InterfaceToString(m["image_source"]),
        LocalGB: schema.InterfaceToString(m["local_gb"]),
        MemoryMB: schema.InterfaceToString(m["memory_mb"]),
        NovaHostID: schema.InterfaceToString(m["nova_host_id"]),
        RootGB: schema.InterfaceToString(m["root_gb"]),
        SwapMB: schema.InterfaceToString(m["swap_mb"]),
        Vcpus: schema.InterfaceToString(m["vcpus"]),
        Capabilities: schema.InterfaceToString(m["capabilities"]),
        
    }
}

// MakeInstanceInfoSlice() makes a slice of InstanceInfo
func MakeInstanceInfoSlice() []*InstanceInfo {
    return []*InstanceInfo{}
}

// InterfaceToInstanceInfoSlice() makes a slice of InstanceInfo
func InterfaceToInstanceInfoSlice(i interface{}) []*InstanceInfo {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*InstanceInfo{}
    for _, item := range list {
        result = append(result, InterfaceToInstanceInfo(item) )
    }
    return result
}



