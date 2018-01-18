package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	Created          string                `json:"created,omitempty"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
}

// String returns json representation of the object
func (model *IpamSubnetType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpamSubnetType makes IpamSubnetType
func MakeIpamSubnetType() *IpamSubnetType {
	return &IpamSubnetType{
		//TODO(nati): Apply default
		AllocUnit: 0,

		AllocationPools: MakeAllocationPoolTypeSlice(),

		DNSServerAddress: MakeIpAddressType(),
		SubnetName:       "",
		AddrFromStart:    false,
		DNSNameservers:   []string{},
		DHCPOptionList:   MakeDhcpOptionsListType(),
		LastModified:     "",
		HostRoutes:       MakeRouteTableType(),
		EnableDHCP:       false,
		DefaultGateway:   MakeIpAddressType(),
		SubnetUUID:       "",
		Subnet:           MakeSubnetType(),
		Created:          "",
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
