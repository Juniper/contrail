package models


// MakeIpamType makes IpamType
func MakeIpamType() *IpamType{
    return &IpamType{
    //TODO(nati): Apply default
    IpamMethod: "",
        IpamDNSMethod: "",
        IpamDNSServer: MakeIpamDnsAddressType(),
        DHCPOptionList: MakeDhcpOptionsListType(),
        HostRoutes: MakeRouteTableType(),
        CidrBlock: MakeSubnetType(),
        
    }
}

// MakeIpamTypeSlice() makes a slice of IpamType
func MakeIpamTypeSlice() []*IpamType {
    return []*IpamType{}
}


