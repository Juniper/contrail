package models

// SecurityLoggingObjectRuleEntryType

import "encoding/json"

// SecurityLoggingObjectRuleEntryType
type SecurityLoggingObjectRuleEntryType struct {
	Rate     int    `json:"rate,omitempty"`
	RuleUUID string `json:"rule_uuid,omitempty"`
}

// String returns json representation of the object
func (model *SecurityLoggingObjectRuleEntryType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryType() *SecurityLoggingObjectRuleEntryType {
	return &SecurityLoggingObjectRuleEntryType{
		//TODO(nati): Apply default
		RuleUUID: "",
		Rate:     0,
	}
}

// MakeSecurityLoggingObjectRuleEntryTypeSlice() makes a slice of SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryTypeSlice() []*SecurityLoggingObjectRuleEntryType {
	return []*SecurityLoggingObjectRuleEntryType{}
}
