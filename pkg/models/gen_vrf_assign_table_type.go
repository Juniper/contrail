package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVrfAssignTableType makes VrfAssignTableType
// nolint
func MakeVrfAssignTableType() *VrfAssignTableType {
	return &VrfAssignTableType{
		//TODO(nati): Apply default

		VRFAssignRule: MakeVrfAssignRuleTypeSlice(),
	}
}

// MakeVrfAssignTableType makes VrfAssignTableType
// nolint
func InterfaceToVrfAssignTableType(i interface{}) *VrfAssignTableType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VrfAssignTableType{
		//TODO(nati): Apply default

		VRFAssignRule: InterfaceToVrfAssignRuleTypeSlice(m["vrf_assign_rule"]),
	}
}

// MakeVrfAssignTableTypeSlice() makes a slice of VrfAssignTableType
// nolint
func MakeVrfAssignTableTypeSlice() []*VrfAssignTableType {
	return []*VrfAssignTableType{}
}

// InterfaceToVrfAssignTableTypeSlice() makes a slice of VrfAssignTableType
// nolint
func InterfaceToVrfAssignTableTypeSlice(i interface{}) []*VrfAssignTableType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VrfAssignTableType{}
	for _, item := range list {
		result = append(result, InterfaceToVrfAssignTableType(item))
	}
	return result
}
