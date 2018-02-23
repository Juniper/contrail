package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakePortMappings makes PortMappings
func MakePortMappings() *PortMappings{
    return &PortMappings{
    //TODO(nati): Apply default
    
            
                PortMappings:  MakePortMapSlice(),
            
        
    }
}

// MakePortMappings makes PortMappings
func InterfaceToPortMappings(i interface{}) *PortMappings{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &PortMappings{
    //TODO(nati): Apply default
    
            
                PortMappings:  InterfaceToPortMapSlice(m["port_mappings"]),
            
        
    }
}

// MakePortMappingsSlice() makes a slice of PortMappings
func MakePortMappingsSlice() []*PortMappings {
    return []*PortMappings{}
}

// InterfaceToPortMappingsSlice() makes a slice of PortMappings
func InterfaceToPortMappingsSlice(i interface{}) []*PortMappings {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*PortMappings{}
    for _, item := range list {
        result = append(result, InterfaceToPortMappings(item) )
    }
    return result
}



