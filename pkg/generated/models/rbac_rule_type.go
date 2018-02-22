package models


// MakeRbacRuleType makes RbacRuleType
func MakeRbacRuleType() *RbacRuleType{
    return &RbacRuleType{
    //TODO(nati): Apply default
    RuleObject: "",
        
            
                RulePerms:  MakeRbacPermTypeSlice(),
            
        RuleField: "",
        
    }
}

// MakeRbacRuleTypeSlice() makes a slice of RbacRuleType
func MakeRbacRuleTypeSlice() []*RbacRuleType {
    return []*RbacRuleType{}
}


