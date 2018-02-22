package models


// MakeIpamSubnetType makes IpamSubnetType
func MakeIpamSubnetType() *IpamSubnetType{
    return &IpamSubnetType{
    //TODO(nati): Apply default
    Subnet: MakeSubnetType(),
        AddrFromStart: false,
        EnableDHCP: false,
        DefaultGateway: "",
        AllocUnit: 0,
        Created: "",
        DNSNameservers: []string{},
        DHCPOptionList: MakeDhcpOptionsListType(),
        SubnetUUID: "",
        
            
                AllocationPools:  MakeAllocationPoolTypeSlice(),
            
        LastModified: "",
        HostRoutes: MakeRouteTableType(),
        DNSServerAddress: "",
        SubnetName: "",
        
    }
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
    return []*IpamSubnetType{}
}


