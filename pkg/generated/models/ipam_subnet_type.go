package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	Created          string                `json:"created,omitempty"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp,omitempty"`
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
		LastModified:  "",
		AddrFromStart: false,
		EnableDHCP:    false,

		AllocationPools: MakeAllocationPoolTypeSlice(),

		HostRoutes:       MakeRouteTableType(),
		DNSServerAddress: MakeIpAddressType(),
		DefaultGateway:   MakeIpAddressType(),
		Created:          "",
		DNSNameservers:   []string{},
		DHCPOptionList:   MakeDhcpOptionsListType(),
		SubnetUUID:       "",
		SubnetName:       "",
		Subnet:           MakeSubnetType(),
		AllocUnit:        0,
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
