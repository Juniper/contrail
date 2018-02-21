package models


// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryType() *SecurityLoggingObjectRuleEntryType{
    return &SecurityLoggingObjectRuleEntryType{
    //TODO(nati): Apply default
    RuleUUID: "",
        Rate: 0,
        
    }
}

// MakeSecurityLoggingObjectRuleEntryTypeSlice() makes a slice of SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryTypeSlice() []*SecurityLoggingObjectRuleEntryType {
    return []*SecurityLoggingObjectRuleEntryType{}
}


