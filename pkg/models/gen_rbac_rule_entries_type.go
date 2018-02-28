package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeRbacRuleEntriesType makes RbacRuleEntriesType
// nolint
func MakeRbacRuleEntriesType() *RbacRuleEntriesType {
	return &RbacRuleEntriesType{
		//TODO(nati): Apply default

		RbacRule: MakeRbacRuleTypeSlice(),
	}
}

// MakeRbacRuleEntriesType makes RbacRuleEntriesType
// nolint
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
// nolint
func MakeRbacRuleEntriesTypeSlice() []*RbacRuleEntriesType {
	return []*RbacRuleEntriesType{}
}

// InterfaceToRbacRuleEntriesTypeSlice() makes a slice of RbacRuleEntriesType
// nolint
func InterfaceToRbacRuleEntriesTypeSlice(i interface{}) []*RbacRuleEntriesType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*RbacRuleEntriesType{}
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleEntriesType(item))
	}
	return result
}
