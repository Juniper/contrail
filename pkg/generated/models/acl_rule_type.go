package models


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

// MakeAclRuleTypeSlice() makes a slice of AclRuleType
func MakeAclRuleTypeSlice() []*AclRuleType {
    return []*AclRuleType{}
}


