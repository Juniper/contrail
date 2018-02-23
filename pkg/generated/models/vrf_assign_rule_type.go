package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeVrfAssignRuleType makes VrfAssignRuleType
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
func InterfaceToVrfAssignRuleType(i interface{}) *VrfAssignRuleType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &VrfAssignRuleType{
		//TODO(nati): Apply default
		RoutingInstance: schema.InterfaceToString(m["routing_instance"]),
		MatchCondition:  InterfaceToMatchConditionType(m["match_condition"]),
		VlanTag:         schema.InterfaceToInt64(m["vlan_tag"]),
		IgnoreACL:       schema.InterfaceToBool(m["ignore_acl"]),
	}
}

// MakeVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
	return []*VrfAssignRuleType{}
}

// InterfaceToVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
func InterfaceToVrfAssignRuleTypeSlice(i interface{}) []*VrfAssignRuleType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*VrfAssignRuleType{}
	for _, item := range list {
		result = append(result, InterfaceToVrfAssignRuleType(item))
	}
	return result
}
