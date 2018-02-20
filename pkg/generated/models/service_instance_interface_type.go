package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeServiceInstanceInterfaceType makes ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceType() *ServiceInstanceInterfaceType{
    return &ServiceInstanceInterfaceType{
    //TODO(nati): Apply default
    VirtualNetwork: "",
        IPAddress: "",
        AllowedAddressPairs: MakeAllowedAddressPairs(),
        StaticRoutes: MakeRouteTableType(),
        
    }
}

// MakeServiceInstanceInterfaceType makes ServiceInstanceInterfaceType
func InterfaceToServiceInstanceInterfaceType(i interface{}) *ServiceInstanceInterfaceType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceInstanceInterfaceType{
    //TODO(nati): Apply default
    VirtualNetwork: schema.InterfaceToString(m["virtual_network"]),
        IPAddress: schema.InterfaceToString(m["ip_address"]),
        AllowedAddressPairs: InterfaceToAllowedAddressPairs(m["allowed_address_pairs"]),
        StaticRoutes: InterfaceToRouteTableType(m["static_routes"]),
        
    }
}

// MakeServiceInstanceInterfaceTypeSlice() makes a slice of ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
    return []*ServiceInstanceInterfaceType{}
}

// InterfaceToServiceInstanceInterfaceTypeSlice() makes a slice of ServiceInstanceInterfaceType
func InterfaceToServiceInstanceInterfaceTypeSlice(i interface{}) []*ServiceInstanceInterfaceType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceInstanceInterfaceType{}
    for _, item := range list {
        result = append(result, InterfaceToServiceInstanceInterfaceType(item) )
    }
    return result
}



