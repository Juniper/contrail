package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeVrfAssignTableType makes VrfAssignTableType
func MakeVrfAssignTableType() *VrfAssignTableType{
    return &VrfAssignTableType{
    //TODO(nati): Apply default
    
            
                VRFAssignRule:  MakeVrfAssignRuleTypeSlice(),
            
        
    }
}

// MakeVrfAssignTableType makes VrfAssignTableType
func InterfaceToVrfAssignTableType(i interface{}) *VrfAssignTableType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &VrfAssignTableType{
    //TODO(nati): Apply default
    
            
                VRFAssignRule:  InterfaceToVrfAssignRuleTypeSlice(m["vrf_assign_rule"]),
            
        
    }
}

// MakeVrfAssignTableTypeSlice() makes a slice of VrfAssignTableType
func MakeVrfAssignTableTypeSlice() []*VrfAssignTableType {
    return []*VrfAssignTableType{}
}

// InterfaceToVrfAssignTableTypeSlice() makes a slice of VrfAssignTableType
func InterfaceToVrfAssignTableTypeSlice(i interface{}) []*VrfAssignTableType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*VrfAssignTableType{}
    for _, item := range list {
        result = append(result, InterfaceToVrfAssignTableType(item) )
    }
    return result
}



