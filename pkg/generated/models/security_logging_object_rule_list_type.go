package models

// SecurityLoggingObjectRuleListType

// SecurityLoggingObjectRuleListType
//proteus:generate
type SecurityLoggingObjectRuleListType struct {
	Rule []*SecurityLoggingObjectRuleEntryType `json:"rule,omitempty"`
}

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType {
	return &SecurityLoggingObjectRuleListType{
		//TODO(nati): Apply default

		Rule: MakeSecurityLoggingObjectRuleEntryTypeSlice(),
	}
}

// MakeSecurityLoggingObjectRuleListTypeSlice() makes a slice of SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
	return []*SecurityLoggingObjectRuleListType{}
}
