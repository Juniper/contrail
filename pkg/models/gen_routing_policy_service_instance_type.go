package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType
// nolint
func MakeRoutingPolicyServiceInstanceType() *RoutingPolicyServiceInstanceType {
	return &RoutingPolicyServiceInstanceType{
		//TODO(nati): Apply default
		RightSequence: "",
		LeftSequence:  "",
	}
}

// MakeRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType
// nolint
func InterfaceToRoutingPolicyServiceInstanceType(i interface{}) *RoutingPolicyServiceInstanceType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RoutingPolicyServiceInstanceType{
		//TODO(nati): Apply default
		RightSequence: common.InterfaceToString(m["right_sequence"]),
		LeftSequence:  common.InterfaceToString(m["left_sequence"]),
	}
}

// MakeRoutingPolicyServiceInstanceTypeSlice() makes a slice of RoutingPolicyServiceInstanceType
// nolint
func MakeRoutingPolicyServiceInstanceTypeSlice() []*RoutingPolicyServiceInstanceType {
	return []*RoutingPolicyServiceInstanceType{}
}

// InterfaceToRoutingPolicyServiceInstanceTypeSlice() makes a slice of RoutingPolicyServiceInstanceType
// nolint
func InterfaceToRoutingPolicyServiceInstanceTypeSlice(i interface{}) []*RoutingPolicyServiceInstanceType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RoutingPolicyServiceInstanceType{}
	for _, item := range list {
		result = append(result, InterfaceToRoutingPolicyServiceInstanceType(item))
	}
	return result
}
