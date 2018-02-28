package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
// nolint
func MakeSecurityLoggingObjectRuleEntryType() *SecurityLoggingObjectRuleEntryType {
	return &SecurityLoggingObjectRuleEntryType{
		//TODO(nati): Apply default
		RuleUUID: "",
		Rate:     0,
	}
}

// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
// nolint
func InterfaceToSecurityLoggingObjectRuleEntryType(i interface{}) *SecurityLoggingObjectRuleEntryType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &SecurityLoggingObjectRuleEntryType{
		//TODO(nati): Apply default
		RuleUUID: common.InterfaceToString(m["rule_uuid"]),
		Rate:     common.InterfaceToInt64(m["rate"]),
	}
}

// MakeSecurityLoggingObjectRuleEntryTypeSlice() makes a slice of SecurityLoggingObjectRuleEntryType
// nolint
func MakeSecurityLoggingObjectRuleEntryTypeSlice() []*SecurityLoggingObjectRuleEntryType {
	return []*SecurityLoggingObjectRuleEntryType{}
}

// InterfaceToSecurityLoggingObjectRuleEntryTypeSlice() makes a slice of SecurityLoggingObjectRuleEntryType
// nolint
func InterfaceToSecurityLoggingObjectRuleEntryTypeSlice(i interface{}) []*SecurityLoggingObjectRuleEntryType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*SecurityLoggingObjectRuleEntryType{}
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObjectRuleEntryType(item))
	}
	return result
}
