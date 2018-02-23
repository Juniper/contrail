package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeRouteType makes RouteType
func MakeRouteType() *RouteType{
    return &RouteType{
    //TODO(nati): Apply default
    Prefix: "",
        NextHop: "",
        CommunityAttributes: MakeCommunityAttributes(),
        NextHopType: "",
        
    }
}

// MakeRouteType makes RouteType
func InterfaceToRouteType(i interface{}) *RouteType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &RouteType{
    //TODO(nati): Apply default
    Prefix: schema.InterfaceToString(m["prefix"]),
        NextHop: schema.InterfaceToString(m["next_hop"]),
        CommunityAttributes: InterfaceToCommunityAttributes(m["community_attributes"]),
        NextHopType: schema.InterfaceToString(m["next_hop_type"]),
        
    }
}

// MakeRouteTypeSlice() makes a slice of RouteType
func MakeRouteTypeSlice() []*RouteType {
    return []*RouteType{}
}

// InterfaceToRouteTypeSlice() makes a slice of RouteType
func InterfaceToRouteTypeSlice(i interface{}) []*RouteType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*RouteType{}
    for _, item := range list {
        result = append(result, InterfaceToRouteType(item) )
    }
    return result
}



