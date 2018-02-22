package models


// MakeAddressType makes AddressType
func MakeAddressType() *AddressType{
    return &AddressType{
    //TODO(nati): Apply default
    SecurityGroup: "",
        Subnet: MakeSubnetType(),
        NetworkPolicy: "",
        
            
                SubnetList:  MakeSubnetTypeSlice(),
            
        VirtualNetwork: "",
        
    }
}

// MakeAddressTypeSlice() makes a slice of AddressType
func MakeAddressTypeSlice() []*AddressType {
    return []*AddressType{}
}


