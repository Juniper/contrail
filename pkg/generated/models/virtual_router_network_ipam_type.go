package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType {
	return &VirtualRouterNetworkIpamType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
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
func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
	return []*VirtualRouterNetworkIpamType{}
}

// InterfaceToVirtualRouterNetworkIpamTypeSlice() makes a slice of VirtualRouterNetworkIpamType
func InterfaceToVirtualRouterNetworkIpamTypeSlice(i interface{}) []*VirtualRouterNetworkIpamType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VirtualRouterNetworkIpamType{}
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouterNetworkIpamType(item))
	}
	return result
}
