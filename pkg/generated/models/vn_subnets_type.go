package models


// MakeVnSubnetsType makes VnSubnetsType
func MakeVnSubnetsType() *VnSubnetsType{
    return &VnSubnetsType{
    //TODO(nati): Apply default
    
            
                IpamSubnets:  MakeIpamSubnetTypeSlice(),
            
        HostRoutes: MakeRouteTableType(),
        
    }
}

// MakeVnSubnetsTypeSlice() makes a slice of VnSubnetsType
func MakeVnSubnetsTypeSlice() []*VnSubnetsType {
    return []*VnSubnetsType{}
}


