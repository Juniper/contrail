package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
// nolint
func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType {
	return &VirtualRouterNetworkIpamType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
// nolint
func InterfaceToVirtualRouterNetworkIpamType(i interface{}) *VirtualRouterNetworkIpamType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VirtualRouterNetworkIpamType{
		//TODO(nati): Apply default

		Subnet: InterfaceToSubnetTypeSlice(m["subnet"]),

		AllocationPools: InterfaceToAllocationPoolTypeSlice(m["allocation_pools"]),
	}
}

// MakeVirtualRouterNetworkIpamTypeSlice() makes a slice of VirtualRouterNetworkIpamType
// nolint
func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
	return []*VirtualRouterNetworkIpamType{}
}

// InterfaceToVirtualRouterNetworkIpamTypeSlice() makes a slice of VirtualRouterNetworkIpamType
// nolint
func InterfaceToVirtualRouterNetworkIpamTypeSlice(i interface{}) []*VirtualRouterNetworkIpamType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualRouterNetworkIpamType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouterNetworkIpamType(item))
	}
	return result
}
