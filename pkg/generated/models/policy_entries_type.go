package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakePolicyEntriesType makes PolicyEntriesType
func MakePolicyEntriesType() *PolicyEntriesType {
	return &PolicyEntriesType{
		//TODO(nati): Apply default

		PolicyRule: MakePolicyRuleTypeSlice(),
	}
}

// MakePolicyEntriesType makes PolicyEntriesType
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
func MakePolicyEntriesTypeSlice() []*PolicyEntriesType {
	return []*PolicyEntriesType{}
}

// InterfaceToPolicyEntriesTypeSlice() makes a slice of PolicyEntriesType
func InterfaceToPolicyEntriesTypeSlice(i interface{}) []*PolicyEntriesType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PolicyEntriesType{}
	for _, item := range list {
		result = append(result, InterfaceToPolicyEntriesType(item))
	}
	return result
}
