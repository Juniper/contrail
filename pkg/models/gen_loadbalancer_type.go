package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerType makes LoadbalancerType
// nolint
func MakeLoadbalancerType() *LoadbalancerType {
	return &LoadbalancerType{
		//TODO(nati): Apply default
		Status:             "",
		ProvisioningStatus: "",
		AdminState:         false,
		VipAddress:         "",
		VipSubnetID:        "",
		OperatingStatus:    "",
	}
}

// MakeLoadbalancerType makes LoadbalancerType
// nolint
func InterfaceToLoadbalancerType(i interface{}) *LoadbalancerType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerType{
		//TODO(nati): Apply default
		Status:             common.InterfaceToString(m["status"]),
		ProvisioningStatus: common.InterfaceToString(m["provisioning_status"]),
		AdminState:         common.InterfaceToBool(m["admin_state"]),
		VipAddress:         common.InterfaceToString(m["vip_address"]),
		VipSubnetID:        common.InterfaceToString(m["vip_subnet_id"]),
		OperatingStatus:    common.InterfaceToString(m["operating_status"]),
	}
}

// MakeLoadbalancerTypeSlice() makes a slice of LoadbalancerType
// nolint
func MakeLoadbalancerTypeSlice() []*LoadbalancerType {
	return []*LoadbalancerType{}
}

// InterfaceToLoadbalancerTypeSlice() makes a slice of LoadbalancerType
// nolint
func InterfaceToLoadbalancerTypeSlice(i interface{}) []*LoadbalancerType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerType{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerType(item))
	}
	return result
}
