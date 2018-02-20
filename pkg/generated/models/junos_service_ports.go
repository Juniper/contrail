package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeJunosServicePorts makes JunosServicePorts
func MakeJunosServicePorts() *JunosServicePorts{
    return &JunosServicePorts{
    //TODO(nati): Apply default
    ServicePort: []string{},
        
    }
}

// MakeJunosServicePorts makes JunosServicePorts
func InterfaceToJunosServicePorts(i interface{}) *JunosServicePorts{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &JunosServicePorts{
    //TODO(nati): Apply default
    ServicePort: schema.InterfaceToStringList(m["service_port"]),
        
    }
}

// MakeJunosServicePortsSlice() makes a slice of JunosServicePorts
func MakeJunosServicePortsSlice() []*JunosServicePorts {
    return []*JunosServicePorts{}
}

// InterfaceToJunosServicePortsSlice() makes a slice of JunosServicePorts
func InterfaceToJunosServicePortsSlice(i interface{}) []*JunosServicePorts {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*JunosServicePorts{}
    for _, item := range list {
        result = append(result, InterfaceToJunosServicePorts(item) )
    }
    return result
}



