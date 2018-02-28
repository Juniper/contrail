package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
// nolint
func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType {
	return &SecurityLoggingObjectRuleListType{
		//TODO(nati): Apply default

		Rule: MakeSecurityLoggingObjectRuleEntryTypeSlice(),
	}
}

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
// nolint
func InterfaceToSecurityLoggingObjectRuleListType(i interface{}) *SecurityLoggingObjectRuleListType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SecurityLoggingObjectRuleListType{
		//TODO(nati): Apply default

		Rule: InterfaceToSecurityLoggingObjectRuleEntryTypeSlice(m["rule"]),
	}
}

// MakeSecurityLoggingObjectRuleListTypeSlice() makes a slice of SecurityLoggingObjectRuleListType
// nolint
func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
	return []*SecurityLoggingObjectRuleListType{}
}

// InterfaceToSecurityLoggingObjectRuleListTypeSlice() makes a slice of SecurityLoggingObjectRuleListType
// nolint
func InterfaceToSecurityLoggingObjectRuleListTypeSlice(i interface{}) []*SecurityLoggingObjectRuleListType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SecurityLoggingObjectRuleListType{}
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObjectRuleListType(item))
	}
	return result
}
