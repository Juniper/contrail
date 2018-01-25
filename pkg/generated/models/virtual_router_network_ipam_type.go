package models

// VirtualRouterNetworkIpamType

// VirtualRouterNetworkIpamType
//proteus:generate
type VirtualRouterNetworkIpamType struct {
	Subnet          []*SubnetType         `json:"subnet,omitempty"`
	AllocationPools []*AllocationPoolType `json:"allocation_pools,omitempty"`
}

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType {
	return &VirtualRouterNetworkIpamType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

// MakeVirtualRouterNetworkIpamTypeSlice() makes a slice of VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
	return []*VirtualRouterNetworkIpamType{}
}
