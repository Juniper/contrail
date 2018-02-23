package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeRbacRuleType makes RbacRuleType
func MakeRbacRuleType() *RbacRuleType {
	return &RbacRuleType{
		//TODO(nati): Apply default
		RuleObject: "",

		RulePerms: MakeRbacPermTypeSlice(),

		RuleField: "",
	}
}

// MakeRbacRuleType makes RbacRuleType
func InterfaceToRbacRuleType(i interface{}) *RbacRuleType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RbacRuleType{
		//TODO(nati): Apply default
		RuleObject: schema.InterfaceToString(m["rule_object"]),

		RulePerms: InterfaceToRbacPermTypeSlice(m["rule_perms"]),

		RuleField: schema.InterfaceToString(m["rule_field"]),
	}
}

// MakeRbacRuleTypeSlice() makes a slice of RbacRuleType
func MakeRbacRuleTypeSlice() []*RbacRuleType {
	return []*RbacRuleType{}
}

// InterfaceToRbacRuleTypeSlice() makes a slice of RbacRuleType
func InterfaceToRbacRuleTypeSlice(i interface{}) []*RbacRuleType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RbacRuleType{}
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleType(item))
	}
	return result
}
