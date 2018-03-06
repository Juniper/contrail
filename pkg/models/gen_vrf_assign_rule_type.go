package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeVrfAssignRuleType makes VrfAssignRuleType
// nolint
func MakeVrfAssignRuleType() *VrfAssignRuleType {
	return &VrfAssignRuleType{
		//TODO(nati): Apply default
		RoutingInstance: "",
		MatchCondition:  MakeMatchConditionType(),
		VlanTag:         0,
		IgnoreACL:       false,
	}
}

// MakeVrfAssignRuleType makes VrfAssignRuleType
// nolint
func InterfaceToVrfAssignRuleType(i interface{}) *VrfAssignRuleType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VrfAssignRuleType{
		//TODO(nati): Apply default
		RoutingInstance: common.InterfaceToString(m["routing_instance"]),
		MatchCondition:  InterfaceToMatchConditionType(m["match_condition"]),
		VlanTag:         common.InterfaceToInt64(m["vlan_tag"]),
		IgnoreACL:       common.InterfaceToBool(m["ignore_acl"]),
	}
}

// MakeVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
// nolint
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
	return []*VrfAssignRuleType{}
}

// InterfaceToVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
// nolint
func InterfaceToVrfAssignRuleTypeSlice(i interface{}) []*VrfAssignRuleType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VrfAssignRuleType{}
	for _, item := range list {
		result = append(result, InterfaceToVrfAssignRuleType(item))
	}
	return result
}
