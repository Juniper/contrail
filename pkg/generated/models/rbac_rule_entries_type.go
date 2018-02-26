package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeRbacRuleEntriesType makes RbacRuleEntriesType
func MakeRbacRuleEntriesType() *RbacRuleEntriesType {
	return &RbacRuleEntriesType{
		//TODO(nati): Apply default

		RbacRule: MakeRbacRuleTypeSlice(),
	}
}

// MakeRbacRuleEntriesType makes RbacRuleEntriesType
func InterfaceToRbacRuleEntriesType(i interface{}) *RbacRuleEntriesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &RbacRuleEntriesType{
		//TODO(nati): Apply default

		RbacRule: InterfaceToRbacRuleTypeSlice(m["rbac_rule"]),
	}
}

// MakeRbacRuleEntriesTypeSlice() makes a slice of RbacRuleEntriesType
func MakeRbacRuleEntriesTypeSlice() []*RbacRuleEntriesType {
	return []*RbacRuleEntriesType{}
}

// InterfaceToRbacRuleEntriesTypeSlice() makes a slice of RbacRuleEntriesType
func InterfaceToRbacRuleEntriesTypeSlice(i interface{}) []*RbacRuleEntriesType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RbacRuleEntriesType{}
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleEntriesType(item))
	}
	return result
}
