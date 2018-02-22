package models


// MakeVrfAssignRuleType makes VrfAssignRuleType
func MakeVrfAssignRuleType() *VrfAssignRuleType{
    return &VrfAssignRuleType{
    //TODO(nati): Apply default
    RoutingInstance: "",
        MatchCondition: MakeMatchConditionType(),
        VlanTag: 0,
        IgnoreACL: false,
        
    }
}

// MakeVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
    return []*VrfAssignRuleType{}
}


