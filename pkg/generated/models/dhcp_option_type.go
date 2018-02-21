package models


// MakeDhcpOptionType makes DhcpOptionType
func MakeDhcpOptionType() *DhcpOptionType{
    return &DhcpOptionType{
    //TODO(nati): Apply default
    DHCPOptionValue: "",
        DHCPOptionValueBytes: "",
        DHCPOptionName: "",
        
    }
}

// MakeDhcpOptionTypeSlice() makes a slice of DhcpOptionType
func MakeDhcpOptionTypeSlice() []*DhcpOptionType {
    return []*DhcpOptionType{}
}


