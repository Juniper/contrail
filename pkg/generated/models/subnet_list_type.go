package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeSubnetListType makes SubnetListType
func MakeSubnetListType() *SubnetListType{
    return &SubnetListType{
    //TODO(nati): Apply default
    
            
                Subnet:  MakeSubnetTypeSlice(),
            
        
    }
}

// MakeSubnetListType makes SubnetListType
func InterfaceToSubnetListType(i interface{}) *SubnetListType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &SubnetListType{
    //TODO(nati): Apply default
    
            
                Subnet:  InterfaceToSubnetTypeSlice(m["subnet"]),
            
        
    }
}

// MakeSubnetListTypeSlice() makes a slice of SubnetListType
func MakeSubnetListTypeSlice() []*SubnetListType {
    return []*SubnetListType{}
}

// InterfaceToSubnetListTypeSlice() makes a slice of SubnetListType
func InterfaceToSubnetListTypeSlice(i interface{}) []*SubnetListType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*SubnetListType{}
    for _, item := range list {
        result = append(result, InterfaceToSubnetListType(item) )
    }
    return result
}



