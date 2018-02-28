package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRbacRuleType makes RbacRuleType
// nolint
func MakeRbacRuleType() *RbacRuleType {
	return &RbacRuleType{
		//TODO(nati): Apply default
		RuleObject: "",

		RulePerms: MakeRbacPermTypeSlice(),

		RuleField: "",
	}
}

// MakeRbacRuleType makes RbacRuleType
// nolint
func InterfaceToRbacRuleType(i interface{}) *RbacRuleType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RbacRuleType{
		//TODO(nati): Apply default
		RuleObject: common.InterfaceToString(m["rule_object"]),

		RulePerms: InterfaceToRbacPermTypeSlice(m["rule_perms"]),

		RuleField: common.InterfaceToString(m["rule_field"]),
	}
}

// MakeRbacRuleTypeSlice() makes a slice of RbacRuleType
// nolint
func MakeRbacRuleTypeSlice() []*RbacRuleType {
	return []*RbacRuleType{}
}

// InterfaceToRbacRuleTypeSlice() makes a slice of RbacRuleType
// nolint
func InterfaceToRbacRuleTypeSlice(i interface{}) []*RbacRuleType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RbacRuleType{}
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleType(item))
	}
	return result
}
