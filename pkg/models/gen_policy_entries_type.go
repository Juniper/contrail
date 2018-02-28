package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePolicyEntriesType makes PolicyEntriesType
// nolint
func MakePolicyEntriesType() *PolicyEntriesType {
	return &PolicyEntriesType{
		//TODO(nati): Apply default

		PolicyRule: MakePolicyRuleTypeSlice(),
	}
}

// MakePolicyEntriesType makes PolicyEntriesType
// nolint
func InterfaceToPolicyEntriesType(i interface{}) *PolicyEntriesType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PolicyEntriesType{
		//TODO(nati): Apply default

		PolicyRule: InterfaceToPolicyRuleTypeSlice(m["policy_rule"]),
	}
}

// MakePolicyEntriesTypeSlice() makes a slice of PolicyEntriesType
// nolint
func MakePolicyEntriesTypeSlice() []*PolicyEntriesType {
	return []*PolicyEntriesType{}
}

// InterfaceToPolicyEntriesTypeSlice() makes a slice of PolicyEntriesType
// nolint
func InterfaceToPolicyEntriesTypeSlice(i interface{}) []*PolicyEntriesType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PolicyEntriesType{}
	for _, item := range list {
		result = append(result, InterfaceToPolicyEntriesType(item))
	}
	return result
}
