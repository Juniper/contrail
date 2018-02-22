package models


// MakeIpamSubnets makes IpamSubnets
func MakeIpamSubnets() *IpamSubnets{
    return &IpamSubnets{
    //TODO(nati): Apply default
    
            
                Subnets:  MakeIpamSubnetTypeSlice(),
            
        
    }
}

// MakeIpamSubnetsSlice() makes a slice of IpamSubnets
func MakeIpamSubnetsSlice() []*IpamSubnets {
    return []*IpamSubnets{}
}


