package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceType() *RoutingPolicyServiceInstanceType{
    return &RoutingPolicyServiceInstanceType{
    //TODO(nati): Apply default
    RightSequence: "",
        LeftSequence: "",
        
    }
}

// MakeRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType
func InterfaceToRoutingPolicyServiceInstanceType(i interface{}) *RoutingPolicyServiceInstanceType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &RoutingPolicyServiceInstanceType{
    //TODO(nati): Apply default
    RightSequence: schema.InterfaceToString(m["right_sequence"]),
        LeftSequence: schema.InterfaceToString(m["left_sequence"]),
        
    }
}

// MakeRoutingPolicyServiceInstanceTypeSlice() makes a slice of RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceTypeSlice() []*RoutingPolicyServiceInstanceType {
    return []*RoutingPolicyServiceInstanceType{}
}

// InterfaceToRoutingPolicyServiceInstanceTypeSlice() makes a slice of RoutingPolicyServiceInstanceType
func InterfaceToRoutingPolicyServiceInstanceTypeSlice(i interface{}) []*RoutingPolicyServiceInstanceType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*RoutingPolicyServiceInstanceType{}
    for _, item := range list {
        result = append(result, InterfaceToRoutingPolicyServiceInstanceType(item) )
    }
    return result
}



