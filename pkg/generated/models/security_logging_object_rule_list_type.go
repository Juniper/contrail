package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType {
	return &SecurityLoggingObjectRuleListType{
		//TODO(nati): Apply default

		Rule: MakeSecurityLoggingObjectRuleEntryTypeSlice(),
	}
}

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
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
func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
	return []*SecurityLoggingObjectRuleListType{}
}

// InterfaceToSecurityLoggingObjectRuleListTypeSlice() makes a slice of SecurityLoggingObjectRuleListType
func InterfaceToSecurityLoggingObjectRuleListTypeSlice(i interface{}) []*SecurityLoggingObjectRuleListType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SecurityLoggingObjectRuleListType{}
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObjectRuleListType(item))
	}
	return result
}
