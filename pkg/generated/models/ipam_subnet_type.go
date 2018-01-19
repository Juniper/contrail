package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	Created          string                `json:"created,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
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
		LastModified: "",
		HostRoutes:   MakeRouteTableType(),
		Created:      "",

		AllocationPools: MakeAllocationPoolTypeSlice(),

		SubnetUUID:       "",
		DNSServerAddress: MakeIpAddressType(),
		Subnet:           MakeSubnetType(),
		DHCPOptionList:   MakeDhcpOptionsListType(),
		DefaultGateway:   MakeIpAddressType(),
		SubnetName:       "",
		AddrFromStart:    false,
		EnableDHCP:       false,
		AllocUnit:        0,
		DNSNameservers:   []string{},
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
