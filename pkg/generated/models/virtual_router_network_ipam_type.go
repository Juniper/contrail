package models

// VirtualRouterNetworkIpamType

import "encoding/json"

// VirtualRouterNetworkIpamType
type VirtualRouterNetworkIpamType struct {
	Subnet          []*SubnetType         `json:"subnet"`
	AllocationPools []*AllocationPoolType `json:"allocation_pools"`
}

// String returns json representation of the object
func (model *VirtualRouterNetworkIpamType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType {
	return &VirtualRouterNetworkIpamType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

// InterfaceToVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType from interface
func InterfaceToVirtualRouterNetworkIpamType(iData interface{}) *VirtualRouterNetworkIpamType {
	data := iData.(map[string]interface{})
	return &VirtualRouterNetworkIpamType{

		Subnet: InterfaceToSubnetTypeSlice(data["subnet"]),

		//{"description":"List of ip prefix and length for vrouter specific subnets","type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}}

		AllocationPools: InterfaceToAllocationPoolTypeSlice(data["allocation_pools"]),

		//{"description":"List of ranges of ip address for vrouter specific allocation","type":"array","item":{"type":"object","properties":{"end":{"type":"string"},"start":{"type":"string"},"vrouter_specific_pool":{"type":"boolean"}}}}

	}
}

// InterfaceToVirtualRouterNetworkIpamTypeSlice makes a slice of VirtualRouterNetworkIpamType from interface
func InterfaceToVirtualRouterNetworkIpamTypeSlice(data interface{}) []*VirtualRouterNetworkIpamType {
	list := data.([]interface{})
	result := MakeVirtualRouterNetworkIpamTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouterNetworkIpamType(item))
	}
	return result
}

// MakeVirtualRouterNetworkIpamTypeSlice() makes a slice of VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
	return []*VirtualRouterNetworkIpamType{}
}
