package models


// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetType() *FloatingIpPoolSubnetType{
    return &FloatingIpPoolSubnetType{
    //TODO(nati): Apply default
    SubnetUUID: []string{},
        
    }
}

// MakeFloatingIpPoolSubnetTypeSlice() makes a slice of FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetTypeSlice() []*FloatingIpPoolSubnetType {
    return []*FloatingIpPoolSubnetType{}
}


