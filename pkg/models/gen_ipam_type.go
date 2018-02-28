package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeIpamType makes IpamType
// nolint
func MakeIpamType() *IpamType {
	return &IpamType{
		//TODO(nati): Apply default
		IpamMethod:     "",
		IpamDNSMethod:  "",
		IpamDNSServer:  MakeIpamDnsAddressType(),
		DHCPOptionList: MakeDhcpOptionsListType(),
		HostRoutes:     MakeRouteTableType(),
		CidrBlock:      MakeSubnetType(),
	}
}

// MakeIpamType makes IpamType
// nolint
func InterfaceToIpamType(i interface{}) *IpamType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpamType{
		//TODO(nati): Apply default
		IpamMethod:     common.InterfaceToString(m["ipam_method"]),
		IpamDNSMethod:  common.InterfaceToString(m["ipam_dns_method"]),
		IpamDNSServer:  InterfaceToIpamDnsAddressType(m["ipam_dns_server"]),
		DHCPOptionList: InterfaceToDhcpOptionsListType(m["dhcp_option_list"]),
		HostRoutes:     InterfaceToRouteTableType(m["host_routes"]),
		CidrBlock:      InterfaceToSubnetType(m["cidr_block"]),
	}
}

// MakeIpamTypeSlice() makes a slice of IpamType
// nolint
func MakeIpamTypeSlice() []*IpamType {
	return []*IpamType{}
}

// InterfaceToIpamTypeSlice() makes a slice of IpamType
// nolint
func InterfaceToIpamTypeSlice(i interface{}) []*IpamType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpamType{}
	for _, item := range list {
		result = append(result, InterfaceToIpamType(item))
	}
	return result
}
