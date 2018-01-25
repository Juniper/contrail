package models
// IpamSubnetType



import "encoding/json"

// IpamSubnetType 
//proteus:generate
type IpamSubnetType struct {

    Subnet *SubnetType `json:"subnet,omitempty"`
    AddrFromStart bool `json:"addr_from_start"`
    EnableDHCP bool `json:"enable_dhcp"`
    DefaultGateway IpAddressType `json:"default_gateway,omitempty"`
    AllocUnit int `json:"alloc_unit,omitempty"`
    Created string `json:"created,omitempty"`
    DNSNameservers []string `json:"dns_nameservers,omitempty"`
    DHCPOptionList *DhcpOptionsListType `json:"dhcp_option_list,omitempty"`
    SubnetUUID string `json:"subnet_uuid,omitempty"`
    AllocationPools []*AllocationPoolType `json:"allocation_pools,omitempty"`
    LastModified string `json:"last_modified,omitempty"`
    HostRoutes *RouteTableType `json:"host_routes,omitempty"`
    DNSServerAddress IpAddressType `json:"dns_server_address,omitempty"`
    SubnetName string `json:"subnet_name,omitempty"`


}



// String returns json representation of the object
func (model *IpamSubnetType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpamSubnetType makes IpamSubnetType
func MakeIpamSubnetType() *IpamSubnetType{
    return &IpamSubnetType{
    //TODO(nati): Apply default
    Subnet: MakeSubnetType(),
        AddrFromStart: false,
        EnableDHCP: false,
        DefaultGateway: MakeIpAddressType(),
        AllocUnit: 0,
        Created: "",
        DNSNameservers: []string{},
        DHCPOptionList: MakeDhcpOptionsListType(),
        SubnetUUID: "",
        
            
                AllocationPools:  MakeAllocationPoolTypeSlice(),
            
        LastModified: "",
        HostRoutes: MakeRouteTableType(),
        DNSServerAddress: MakeIpAddressType(),
        SubnetName: "",
        
    }
}



// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
    return []*IpamSubnetType{}
}
