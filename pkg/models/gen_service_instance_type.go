package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceInstanceType makes ServiceInstanceType
// nolint
func MakeServiceInstanceType() *ServiceInstanceType {
	return &ServiceInstanceType{
		//TODO(nati): Apply default
		RightVirtualNetwork:      "",
		RightIPAddress:           "",
		AvailabilityZone:         "",
		ManagementVirtualNetwork: "",
		ScaleOut:                 MakeServiceScaleOutType(),
		HaMode:                   "",
		VirtualRouterID:          "",

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		LeftIPAddress:      "",
		LeftVirtualNetwork: "",
		AutoPolicy:         false,
	}
}

// MakeServiceInstanceType makes ServiceInstanceType
// nolint
func InterfaceToServiceInstanceType(i interface{}) *ServiceInstanceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceInstanceType{
		//TODO(nati): Apply default
		RightVirtualNetwork:      common.InterfaceToString(m["right_virtual_network"]),
		RightIPAddress:           common.InterfaceToString(m["right_ip_address"]),
		AvailabilityZone:         common.InterfaceToString(m["availability_zone"]),
		ManagementVirtualNetwork: common.InterfaceToString(m["management_virtual_network"]),
		ScaleOut:                 InterfaceToServiceScaleOutType(m["scale_out"]),
		HaMode:                   common.InterfaceToString(m["ha_mode"]),
		VirtualRouterID:          common.InterfaceToString(m["virtual_router_id"]),

		InterfaceList: InterfaceToServiceInstanceInterfaceTypeSlice(m["interface_list"]),

		LeftIPAddress:      common.InterfaceToString(m["left_ip_address"]),
		LeftVirtualNetwork: common.InterfaceToString(m["left_virtual_network"]),
		AutoPolicy:         common.InterfaceToBool(m["auto_policy"]),
	}
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
// nolint
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
	return []*ServiceInstanceType{}
}

// InterfaceToServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
// nolint
func InterfaceToServiceInstanceTypeSlice(i interface{}) []*ServiceInstanceType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceInstanceType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceInstanceType(item))
	}
	return result
}
