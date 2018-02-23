package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeBaremetalProperties makes BaremetalProperties
func MakeBaremetalProperties() *BaremetalProperties{
    return &BaremetalProperties{
    //TODO(nati): Apply default
    CPUCount: 0,
        CPUArch: "",
        DiskGB: 0,
        MemoryMB: 0,
        
    }
}

// MakeBaremetalProperties makes BaremetalProperties
func InterfaceToBaremetalProperties(i interface{}) *BaremetalProperties{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &BaremetalProperties{
    //TODO(nati): Apply default
    CPUCount: schema.InterfaceToInt64(m["cpu_count"]),
        CPUArch: schema.InterfaceToString(m["cpu_arch"]),
        DiskGB: schema.InterfaceToInt64(m["disk_gb"]),
        MemoryMB: schema.InterfaceToInt64(m["memory_mb"]),
        
    }
}

// MakeBaremetalPropertiesSlice() makes a slice of BaremetalProperties
func MakeBaremetalPropertiesSlice() []*BaremetalProperties {
    return []*BaremetalProperties{}
}

// InterfaceToBaremetalPropertiesSlice() makes a slice of BaremetalProperties
func InterfaceToBaremetalPropertiesSlice(i interface{}) []*BaremetalProperties {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*BaremetalProperties{}
    for _, item := range list {
        result = append(result, InterfaceToBaremetalProperties(item) )
    }
    return result
}



