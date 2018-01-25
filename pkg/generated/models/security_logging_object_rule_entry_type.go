package models

// SecurityLoggingObjectRuleEntryType

// SecurityLoggingObjectRuleEntryType
//proteus:generate
type SecurityLoggingObjectRuleEntryType struct {
	RuleUUID string `json:"rule_uuid,omitempty"`
	Rate     int    `json:"rate,omitempty"`
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
