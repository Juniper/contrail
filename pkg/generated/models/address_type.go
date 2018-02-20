package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAddressType makes AddressType
func MakeAddressType() *AddressType{
    return &AddressType{
    //TODO(nati): Apply default
    SecurityGroup: "",
        Subnet: MakeSubnetType(),
        NetworkPolicy: "",
        
            
                SubnetList:  MakeSubnetTypeSlice(),
            
        VirtualNetwork: "",
        
    }
}

// MakeAddressType makes AddressType
func InterfaceToAddressType(i interface{}) *AddressType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AddressType{
    //TODO(nati): Apply default
    SecurityGroup: schema.InterfaceToString(m["security_group"]),
        Subnet: InterfaceToSubnetType(m["subnet"]),
        NetworkPolicy: schema.InterfaceToString(m["network_policy"]),
        
            
                SubnetList:  InterfaceToSubnetTypeSlice(m["subnet_list"]),
            
        VirtualNetwork: schema.InterfaceToString(m["virtual_network"]),
        
    }
}

// MakeAddressTypeSlice() makes a slice of AddressType
func MakeAddressTypeSlice() []*AddressType {
    return []*AddressType{}
}

// InterfaceToAddressTypeSlice() makes a slice of AddressType
func InterfaceToAddressTypeSlice(i interface{}) []*AddressType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AddressType{}
    for _, item := range list {
        result = append(result, InterfaceToAddressType(item) )
    }
    return result
}



