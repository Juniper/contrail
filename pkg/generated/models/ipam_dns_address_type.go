package models


// MakeIpamDnsAddressType makes IpamDnsAddressType
func MakeIpamDnsAddressType() *IpamDnsAddressType{
    return &IpamDnsAddressType{
    //TODO(nati): Apply default
    TenantDNSServerAddress: MakeIpAddressesType(),
        VirtualDNSServerName: "",
        
    }
}

// MakeIpamDnsAddressTypeSlice() makes a slice of IpamDnsAddressType
func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
    return []*IpamDnsAddressType{}
}


