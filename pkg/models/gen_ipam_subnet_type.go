package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeIpamSubnetType makes IpamSubnetType
// nolint
func MakeIpamSubnetType() *IpamSubnetType {
	return &IpamSubnetType{
		//TODO(nati): Apply default
		Subnet:         MakeSubnetType(),
		AddrFromStart:  false,
		EnableDHCP:     false,
		DefaultGateway: "",
		AllocUnit:      0,
		Created:        "",
		DNSNameservers: []string{},
		DHCPOptionList: MakeDhcpOptionsListType(),
		SubnetUUID:     "",

		AllocationPools: MakeAllocationPoolTypeSlice(),

		LastModified:     "",
		HostRoutes:       MakeRouteTableType(),
		DNSServerAddress: "",
		SubnetName:       "",
	}
}

// MakeIpamSubnetType makes IpamSubnetType
// nolint
func InterfaceToIpamSubnetType(i interface{}) *IpamSubnetType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpamSubnetType{
		//TODO(nati): Apply default
		Subnet:         InterfaceToSubnetType(m["subnet"]),
		AddrFromStart:  common.InterfaceToBool(m["addr_from_start"]),
		EnableDHCP:     common.InterfaceToBool(m["enable_dhcp"]),
		DefaultGateway: common.InterfaceToString(m["default_gateway"]),
		AllocUnit:      common.InterfaceToInt64(m["alloc_unit"]),
		Created:        common.InterfaceToString(m["created"]),
		DNSNameservers: common.InterfaceToStringList(m["dns_nameservers"]),
		DHCPOptionList: InterfaceToDhcpOptionsListType(m["dhcp_option_list"]),
		SubnetUUID:     common.InterfaceToString(m["subnet_uuid"]),

		AllocationPools: InterfaceToAllocationPoolTypeSlice(m["allocation_pools"]),

		LastModified:     common.InterfaceToString(m["last_modified"]),
		HostRoutes:       InterfaceToRouteTableType(m["host_routes"]),
		DNSServerAddress: common.InterfaceToString(m["dns_server_address"]),
		SubnetName:       common.InterfaceToString(m["subnet_name"]),
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
// nolint
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}

// InterfaceToIpamSubnetTypeSlice() makes a slice of IpamSubnetType
// nolint
func InterfaceToIpamSubnetTypeSlice(i interface{}) []*IpamSubnetType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpamSubnetType{}
	for _, item := range list {
		result = append(result, InterfaceToIpamSubnetType(item))
	}
	return result
}
