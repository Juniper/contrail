package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeRouteTableType makes RouteTableType
func MakeRouteTableType() *RouteTableType{
    return &RouteTableType{
    //TODO(nati): Apply default
    
            
                Route:  MakeRouteTypeSlice(),
            
        
    }
}

// MakeRouteTableType makes RouteTableType
func InterfaceToRouteTableType(i interface{}) *RouteTableType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &RouteTableType{
    //TODO(nati): Apply default
    
            
                Route:  InterfaceToRouteTypeSlice(m["route"]),
            
        
    }
}

// MakeRouteTableTypeSlice() makes a slice of RouteTableType
func MakeRouteTableTypeSlice() []*RouteTableType {
    return []*RouteTableType{}
}

// InterfaceToRouteTableTypeSlice() makes a slice of RouteTableType
func InterfaceToRouteTableTypeSlice(i interface{}) []*RouteTableType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*RouteTableType{}
    for _, item := range list {
        result = append(result, InterfaceToRouteTableType(item) )
    }
    return result
}



