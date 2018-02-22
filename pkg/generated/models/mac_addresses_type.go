package models


// MakeMacAddressesType makes MacAddressesType
func MakeMacAddressesType() *MacAddressesType{
    return &MacAddressesType{
    //TODO(nati): Apply default
    MacAddress: []string{},
        
    }
}

// MakeMacAddressesTypeSlice() makes a slice of MacAddressesType
func MakeMacAddressesTypeSlice() []*MacAddressesType {
    return []*MacAddressesType{}
}


