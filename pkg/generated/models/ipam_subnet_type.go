package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	AllocUnit        int                   `json:"alloc_unit"`
	Created          string                `json:"created"`
	DNSNameservers   []string              `json:"dns_nameservers"`
	SubnetUUID       string                `json:"subnet_uuid"`
	SubnetName       string                `json:"subnet_name"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DefaultGateway   IpAddressType         `json:"default_gateway"`
	HostRoutes       *RouteTableType       `json:"host_routes"`
	DNSServerAddress IpAddressType         `json:"dns_server_address"`
	Subnet           *SubnetType           `json:"subnet"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools"`
	LastModified     string                `json:"last_modified"`
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
		DefaultGateway:   MakeIpAddressType(),
		HostRoutes:       MakeRouteTableType(),
		DNSServerAddress: MakeIpAddressType(),
		AddrFromStart:    false,
		DHCPOptionList:   MakeDhcpOptionsListType(),

		AllocationPools: MakeAllocationPoolTypeSlice(),

		LastModified:   "",
		Subnet:         MakeSubnetType(),
		Created:        "",
		AllocUnit:      0,
		SubnetUUID:     "",
		SubnetName:     "",
		DNSNameservers: []string{},
	}
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
