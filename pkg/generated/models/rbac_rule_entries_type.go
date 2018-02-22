package models


// MakeRbacRuleEntriesType makes RbacRuleEntriesType
func MakeRbacRuleEntriesType() *RbacRuleEntriesType{
    return &RbacRuleEntriesType{
    //TODO(nati): Apply default
    
            
                RbacRule:  MakeRbacRuleTypeSlice(),
            
        
    }
}

// MakeRbacRuleEntriesTypeSlice() makes a slice of RbacRuleEntriesType
func MakeRbacRuleEntriesTypeSlice() []*RbacRuleEntriesType {
    return []*RbacRuleEntriesType{}
}


