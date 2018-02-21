package models


// MakeAllowedAddressPair makes AllowedAddressPair
func MakeAllowedAddressPair() *AllowedAddressPair{
    return &AllowedAddressPair{
    //TODO(nati): Apply default
    IP: MakeSubnetType(),
        Mac: "",
        AddressMode: "",
        
    }
}

// MakeAllowedAddressPairSlice() makes a slice of AllowedAddressPair
func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
    return []*AllowedAddressPair{}
}


