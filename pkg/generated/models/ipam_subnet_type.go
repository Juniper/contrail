package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	Created          string                `json:"created,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
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
		SubnetUUID:     "",
		DHCPOptionList: MakeDhcpOptionsListType(),

		AllocationPools: MakeAllocationPoolTypeSlice(),

		HostRoutes:       MakeRouteTableType(),
		AllocUnit:        0,
		DNSNameservers:   []string{},
		DNSServerAddress: MakeIpAddressType(),
		Subnet:           MakeSubnetType(),
		EnableDHCP:       false,
		DefaultGateway:   MakeIpAddressType(),
		SubnetName:       "",
		AddrFromStart:    false,
		Created:          "",
		LastModified:     "",
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
