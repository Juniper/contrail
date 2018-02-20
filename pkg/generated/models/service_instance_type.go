package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeServiceInstanceType makes ServiceInstanceType
func MakeServiceInstanceType() *ServiceInstanceType{
    return &ServiceInstanceType{
    //TODO(nati): Apply default
    RightVirtualNetwork: "",
        RightIPAddress: "",
        AvailabilityZone: "",
        ManagementVirtualNetwork: "",
        ScaleOut: MakeServiceScaleOutType(),
        HaMode: "",
        VirtualRouterID: "",
        
            
                InterfaceList:  MakeServiceInstanceInterfaceTypeSlice(),
            
        LeftIPAddress: "",
        LeftVirtualNetwork: "",
        AutoPolicy: false,
        
    }
}

// MakeServiceInstanceType makes ServiceInstanceType
func InterfaceToServiceInstanceType(i interface{}) *ServiceInstanceType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &ServiceInstanceType{
    //TODO(nati): Apply default
    RightVirtualNetwork: schema.InterfaceToString(m["right_virtual_network"]),
        RightIPAddress: schema.InterfaceToString(m["right_ip_address"]),
        AvailabilityZone: schema.InterfaceToString(m["availability_zone"]),
        ManagementVirtualNetwork: schema.InterfaceToString(m["management_virtual_network"]),
        ScaleOut: InterfaceToServiceScaleOutType(m["scale_out"]),
        HaMode: schema.InterfaceToString(m["ha_mode"]),
        VirtualRouterID: schema.InterfaceToString(m["virtual_router_id"]),
        
            
                InterfaceList:  InterfaceToServiceInstanceInterfaceTypeSlice(m["interface_list"]),
            
        LeftIPAddress: schema.InterfaceToString(m["left_ip_address"]),
        LeftVirtualNetwork: schema.InterfaceToString(m["left_virtual_network"]),
        AutoPolicy: schema.InterfaceToBool(m["auto_policy"]),
        
    }
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
    return []*ServiceInstanceType{}
}

// InterfaceToServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func InterfaceToServiceInstanceTypeSlice(i interface{}) []*ServiceInstanceType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*ServiceInstanceType{}
    for _, item := range list {
        result = append(result, InterfaceToServiceInstanceType(item) )
    }
    return result
}



