package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	Created          string                `json:"created,omitempty"`
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
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
		Subnet:           MakeSubnetType(),
		EnableDHCP:       false,
		AllocUnit:        0,
		DNSNameservers:   []string{},
		HostRoutes:       MakeRouteTableType(),
		AddrFromStart:    false,
		DNSServerAddress: MakeIpAddressType(),
		SubnetName:       "",
		DefaultGateway:   MakeIpAddressType(),
		Created:          "",
		SubnetUUID:       "",
		LastModified:     "",
		DHCPOptionList:   MakeDhcpOptionsListType(),

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
