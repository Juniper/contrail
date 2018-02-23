package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesType() *EncapsulationPrioritiesType{
    return &EncapsulationPrioritiesType{
    //TODO(nati): Apply default
    
            
        
    }
}

// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
func InterfaceToEncapsulationPrioritiesType(i interface{}) *EncapsulationPrioritiesType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &EncapsulationPrioritiesType{
    //TODO(nati): Apply default
    
            
        
    }
}

// MakeEncapsulationPrioritiesTypeSlice() makes a slice of EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesTypeSlice() []*EncapsulationPrioritiesType {
    return []*EncapsulationPrioritiesType{}
}

// InterfaceToEncapsulationPrioritiesTypeSlice() makes a slice of EncapsulationPrioritiesType
func InterfaceToEncapsulationPrioritiesTypeSlice(i interface{}) []*EncapsulationPrioritiesType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*EncapsulationPrioritiesType{}
    for _, item := range list {
        result = append(result, InterfaceToEncapsulationPrioritiesType(item) )
    }
    return result
}



