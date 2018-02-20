package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeIpamType makes IpamType
func MakeIpamType() *IpamType{
    return &IpamType{
    //TODO(nati): Apply default
    IpamMethod: "",
        IpamDNSMethod: "",
        IpamDNSServer: MakeIpamDnsAddressType(),
        DHCPOptionList: MakeDhcpOptionsListType(),
        HostRoutes: MakeRouteTableType(),
        CidrBlock: MakeSubnetType(),
        
    }
}

// MakeIpamType makes IpamType
func InterfaceToIpamType(i interface{}) *IpamType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &IpamType{
    //TODO(nati): Apply default
    IpamMethod: schema.InterfaceToString(m["ipam_method"]),
        IpamDNSMethod: schema.InterfaceToString(m["ipam_dns_method"]),
        IpamDNSServer: InterfaceToIpamDnsAddressType(m["ipam_dns_server"]),
        DHCPOptionList: InterfaceToDhcpOptionsListType(m["dhcp_option_list"]),
        HostRoutes: InterfaceToRouteTableType(m["host_routes"]),
        CidrBlock: InterfaceToSubnetType(m["cidr_block"]),
        
    }
}

// MakeIpamTypeSlice() makes a slice of IpamType
func MakeIpamTypeSlice() []*IpamType {
    return []*IpamType{}
}

// InterfaceToIpamTypeSlice() makes a slice of IpamType
func InterfaceToIpamTypeSlice(i interface{}) []*IpamType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*IpamType{}
    for _, item := range list {
        result = append(result, InterfaceToIpamType(item) )
    }
    return result
}



