package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLoadbalancerType makes LoadbalancerType
func MakeLoadbalancerType() *LoadbalancerType{
    return &LoadbalancerType{
    //TODO(nati): Apply default
    Status: "",
        ProvisioningStatus: "",
        AdminState: false,
        VipAddress: "",
        VipSubnetID: "",
        OperatingStatus: "",
        
    }
}

// MakeLoadbalancerType makes LoadbalancerType
func InterfaceToLoadbalancerType(i interface{}) *LoadbalancerType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &LoadbalancerType{
    //TODO(nati): Apply default
    Status: schema.InterfaceToString(m["status"]),
        ProvisioningStatus: schema.InterfaceToString(m["provisioning_status"]),
        AdminState: schema.InterfaceToBool(m["admin_state"]),
        VipAddress: schema.InterfaceToString(m["vip_address"]),
        VipSubnetID: schema.InterfaceToString(m["vip_subnet_id"]),
        OperatingStatus: schema.InterfaceToString(m["operating_status"]),
        
    }
}

// MakeLoadbalancerTypeSlice() makes a slice of LoadbalancerType
func MakeLoadbalancerTypeSlice() []*LoadbalancerType {
    return []*LoadbalancerType{}
}

// InterfaceToLoadbalancerTypeSlice() makes a slice of LoadbalancerType
func InterfaceToLoadbalancerTypeSlice(i interface{}) []*LoadbalancerType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*LoadbalancerType{}
    for _, item := range list {
        result = append(result, InterfaceToLoadbalancerType(item) )
    }
    return result
}



