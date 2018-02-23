package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeAclRuleType makes AclRuleType
func MakeAclRuleType() *AclRuleType{
    return &AclRuleType{
    //TODO(nati): Apply default
    RuleUUID: "",
        MatchCondition: MakeMatchConditionType(),
        Direction: "",
        ActionList: MakeActionListType(),
        
    }
}

// MakeAclRuleType makes AclRuleType
func InterfaceToAclRuleType(i interface{}) *AclRuleType{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &AclRuleType{
    //TODO(nati): Apply default
    RuleUUID: schema.InterfaceToString(m["rule_uuid"]),
        MatchCondition: InterfaceToMatchConditionType(m["match_condition"]),
        Direction: schema.InterfaceToString(m["direction"]),
        ActionList: InterfaceToActionListType(m["action_list"]),
        
    }
}

// MakeAclRuleTypeSlice() makes a slice of AclRuleType
func MakeAclRuleTypeSlice() []*AclRuleType {
    return []*AclRuleType{}
}

// InterfaceToAclRuleTypeSlice() makes a slice of AclRuleType
func InterfaceToAclRuleTypeSlice(i interface{}) []*AclRuleType {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*AclRuleType{}
    for _, item := range list {
        result = append(result, InterfaceToAclRuleType(item) )
    }
    return result
}



