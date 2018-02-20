package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentType() *DiscoveryServiceAssignmentType{
    return &DiscoveryServiceAssignmentType{
    //TODO(nati): Apply default
    
            
                Subscriber:  MakeDiscoveryPubSubEndPointTypeSlice(),
            
        Publisher: MakeDiscoveryPubSubEndPointType(),
        
    }
}

// MakeDiscoveryServiceAssignmentType makes DiscoveryServiceAssignmentType
func InterfaceToDiscoveryServiceAssignmentType(i interface{}) *DiscoveryServiceAssignmentType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &DiscoveryServiceAssignmentType{
    //TODO(nati): Apply default
    
            
                Subscriber:  InterfaceToDiscoveryPubSubEndPointTypeSlice(m["subscriber"]),
            
        Publisher: InterfaceToDiscoveryPubSubEndPointType(m["publisher"]),
        
    }
}

// MakeDiscoveryServiceAssignmentTypeSlice() makes a slice of DiscoveryServiceAssignmentType
func MakeDiscoveryServiceAssignmentTypeSlice() []*DiscoveryServiceAssignmentType {
    return []*DiscoveryServiceAssignmentType{}
}

// InterfaceToDiscoveryServiceAssignmentTypeSlice() makes a slice of DiscoveryServiceAssignmentType
func InterfaceToDiscoveryServiceAssignmentTypeSlice(i interface{}) []*DiscoveryServiceAssignmentType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*DiscoveryServiceAssignmentType{}
    for _, item := range list {
        result = append(result, InterfaceToDiscoveryServiceAssignmentType(item) )
    }
    return result
}



