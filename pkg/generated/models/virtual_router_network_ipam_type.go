package models
// VirtualRouterNetworkIpamType



import "encoding/json"

// VirtualRouterNetworkIpamType 
//proteus:generate
type VirtualRouterNetworkIpamType struct {

    Subnet []*SubnetType `json:"subnet,omitempty"`
    AllocationPools []*AllocationPoolType `json:"allocation_pools,omitempty"`


}



// String returns json representation of the object
func (model *VirtualRouterNetworkIpamType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

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
