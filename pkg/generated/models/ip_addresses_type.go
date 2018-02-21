package models


// MakeIpAddressesType makes IpAddressesType
func MakeIpAddressesType() *IpAddressesType{
    return &IpAddressesType{
    //TODO(nati): Apply default
    IPAddress: "",
        
    }
}

// MakeIpAddressesTypeSlice() makes a slice of IpAddressesType
func MakeIpAddressesTypeSlice() []*IpAddressesType {
    return []*IpAddressesType{}
}


