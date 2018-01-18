package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	Created          string                `json:"created,omitempty"`
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
		EnableDHCP:     false,
		DefaultGateway: MakeIpAddressType(),
		DNSNameservers: []string{},
		DHCPOptionList: MakeDhcpOptionsListType(),
		LastModified:   "",
		AddrFromStart:  false,
		Created:        "",
		SubnetUUID:     "",

		AllocationPools: MakeAllocationPoolTypeSlice(),

		SubnetName:       "",
		Subnet:           MakeSubnetType(),
		AllocUnit:        0,
		HostRoutes:       MakeRouteTableType(),
		DNSServerAddress: MakeIpAddressType(),
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
