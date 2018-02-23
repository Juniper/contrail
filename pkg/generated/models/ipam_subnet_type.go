package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeIpamSubnetType makes IpamSubnetType
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
func InterfaceToIpamSubnetType(i interface{}) *IpamSubnetType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &IpamSubnetType{
		//TODO(nati): Apply default
		Subnet:         InterfaceToSubnetType(m["subnet"]),
		AddrFromStart:  schema.InterfaceToBool(m["addr_from_start"]),
		EnableDHCP:     schema.InterfaceToBool(m["enable_dhcp"]),
		DefaultGateway: schema.InterfaceToString(m["default_gateway"]),
		AllocUnit:      schema.InterfaceToInt64(m["alloc_unit"]),
		Created:        schema.InterfaceToString(m["created"]),
		DNSNameservers: schema.InterfaceToStringList(m["dns_nameservers"]),
		DHCPOptionList: InterfaceToDhcpOptionsListType(m["dhcp_option_list"]),
		SubnetUUID:     schema.InterfaceToString(m["subnet_uuid"]),

		AllocationPools: InterfaceToAllocationPoolTypeSlice(m["allocation_pools"]),

		LastModified:     schema.InterfaceToString(m["last_modified"]),
		HostRoutes:       InterfaceToRouteTableType(m["host_routes"]),
		DNSServerAddress: schema.InterfaceToString(m["dns_server_address"]),
		SubnetName:       schema.InterfaceToString(m["subnet_name"]),
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}

// InterfaceToIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func InterfaceToIpamSubnetTypeSlice(i interface{}) []*IpamSubnetType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*IpamSubnetType{}
	for _, item := range list {
		result = append(result, InterfaceToIpamSubnetType(item))
	}
	return result
}
