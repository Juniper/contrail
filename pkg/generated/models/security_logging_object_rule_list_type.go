package models


// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType{
    return &SecurityLoggingObjectRuleListType{
    //TODO(nati): Apply default
    
            
                Rule:  MakeSecurityLoggingObjectRuleEntryTypeSlice(),
            
        
    }
}

// MakeSecurityLoggingObjectRuleListTypeSlice() makes a slice of SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
    return []*SecurityLoggingObjectRuleListType{}
}


