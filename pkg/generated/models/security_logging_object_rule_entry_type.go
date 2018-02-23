package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryType() *SecurityLoggingObjectRuleEntryType{
    return &SecurityLoggingObjectRuleEntryType{
    //TODO(nati): Apply default
    RuleUUID: "",
        Rate: 0,
        
    }
}

// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
func InterfaceToSecurityLoggingObjectRuleEntryType(i interface{}) *SecurityLoggingObjectRuleEntryType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &SecurityLoggingObjectRuleEntryType{
    //TODO(nati): Apply default
    RuleUUID: schema.InterfaceToString(m["rule_uuid"]),
        Rate: schema.InterfaceToInt64(m["rate"]),
        
    }
}

// MakeSecurityLoggingObjectRuleEntryTypeSlice() makes a slice of SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryTypeSlice() []*SecurityLoggingObjectRuleEntryType {
    return []*SecurityLoggingObjectRuleEntryType{}
}

// InterfaceToSecurityLoggingObjectRuleEntryTypeSlice() makes a slice of SecurityLoggingObjectRuleEntryType
func InterfaceToSecurityLoggingObjectRuleEntryTypeSlice(i interface{}) []*SecurityLoggingObjectRuleEntryType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*SecurityLoggingObjectRuleEntryType{}
    for _, item := range list {
        result = append(result, InterfaceToSecurityLoggingObjectRuleEntryType(item) )
    }
    return result
}



