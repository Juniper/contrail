package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	LastModified     string                `json:"last_modified,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
	Created          string                `json:"created,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
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
		EnableDHCP:       false,
		DNSNameservers:   []string{},
		SubnetName:       "",
		AllocUnit:        0,
		DHCPOptionList:   MakeDhcpOptionsListType(),
		DefaultGateway:   MakeIpAddressType(),
		SubnetUUID:       "",
		LastModified:     "",
		HostRoutes:       MakeRouteTableType(),
		DNSServerAddress: MakeIpAddressType(),
		Subnet:           MakeSubnetType(),
		AddrFromStart:    false,
		Created:          "",

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
