package models


// MakeOpenStackAddress makes OpenStackAddress
func MakeOpenStackAddress() *OpenStackAddress{
    return &OpenStackAddress{
    //TODO(nati): Apply default
    Addr: "",
        
    }
}

// MakeOpenStackAddressSlice() makes a slice of OpenStackAddress
func MakeOpenStackAddressSlice() []*OpenStackAddress {
    return []*OpenStackAddress{}
}


