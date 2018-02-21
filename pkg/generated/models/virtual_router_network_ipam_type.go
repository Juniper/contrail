package models


// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType{
    return &VirtualRouterNetworkIpamType{
    //TODO(nati): Apply default
    
            
                Subnet:  MakeSubnetTypeSlice(),
            
        
            
                AllocationPools:  MakeAllocationPoolTypeSlice(),
            
        
    }
}

// MakeVirtualRouterNetworkIpamTypeSlice() makes a slice of VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
    return []*VirtualRouterNetworkIpamType{}
}


