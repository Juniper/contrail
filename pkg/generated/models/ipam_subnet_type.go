package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	Created          string                `json:"created,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
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
		Subnet:           MakeSubnetType(),
		SubnetUUID:       "",
		EnableDHCP:       false,
		AllocUnit:        0,
		Created:          "",
		DNSNameservers:   []string{},
		DHCPOptionList:   MakeDhcpOptionsListType(),
		HostRoutes:       MakeRouteTableType(),
		DNSServerAddress: MakeIpAddressType(),
		SubnetName:       "",
		AddrFromStart:    false,
		DefaultGateway:   MakeIpAddressType(),

		AllocationPools: MakeAllocationPoolTypeSlice(),

		LastModified: "",
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
