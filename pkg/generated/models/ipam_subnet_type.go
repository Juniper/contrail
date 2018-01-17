package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	DefaultGateway   IpAddressType         `json:"default_gateway,omitempty"`
	SubnetUUID       string                `json:"subnet_uuid,omitempty"`
	Subnet           *SubnetType           `json:"subnet,omitempty"`
	EnableDHCP       bool                  `json:"enable_dhcp,omitempty"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list,omitempty"`
	LastModified     string                `json:"last_modified,omitempty"`
	HostRoutes       *RouteTableType       `json:"host_routes,omitempty"`
	SubnetName       string                `json:"subnet_name,omitempty"`
	AddrFromStart    bool                  `json:"addr_from_start,omitempty"`
	Created          string                `json:"created,omitempty"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools,omitempty"`
	DNSServerAddress IpAddressType         `json:"dns_server_address,omitempty"`
	AllocUnit        int                   `json:"alloc_unit,omitempty"`
	DNSNameservers   []string              `json:"dns_nameservers,omitempty"`
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
		Subnet:         MakeSubnetType(),
		EnableDHCP:     false,
		DefaultGateway: MakeIpAddressType(),
		SubnetUUID:     "",
		AddrFromStart:  false,
		Created:        "",
		DHCPOptionList: MakeDhcpOptionsListType(),
		LastModified:   "",
		HostRoutes:     MakeRouteTableType(),
		SubnetName:     "",
		AllocUnit:      0,
		DNSNameservers: []string{},

		AllocationPools: MakeAllocationPoolTypeSlice(),

		DNSServerAddress: MakeIpAddressType(),
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
