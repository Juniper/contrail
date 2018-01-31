package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	LastModified     string                `json:"last_modified,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	Created          string                `json:"created,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
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
		SubnetName:     "",
		Subnet:         MakeSubnetType(),
		AddrFromStart:  false,
		LastModified:   "",
		DHCPOptionList: MakeDhcpOptionsListType(),

		AllocationPools: MakeAllocationPoolTypeSlice(),

		EnableDHCP:       false,
		DefaultGateway:   MakeIpAddressType(),
		DNSNameservers:   []string{},
		Created:          "",
		DNSServerAddress: MakeIpAddressType(),
		AllocUnit:        0,
		SubnetUUID:       "",
		HostRoutes:       MakeRouteTableType(),
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
