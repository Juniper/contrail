package models
// AclRuleType



import "encoding/json"

// AclRuleType 
//proteus:generate
type AclRuleType struct {

    RuleUUID string `json:"rule_uuid,omitempty"`
    MatchCondition *MatchConditionType `json:"match_condition,omitempty"`
    Direction DirectionType `json:"direction,omitempty"`
    ActionList *ActionListType `json:"action_list,omitempty"`


}



// String returns json representation of the object
func (model *AclRuleType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAclRuleType makes AclRuleType
func MakeAclRuleType() *AclRuleType{
    return &AclRuleType{
    //TODO(nati): Apply default
    RuleUUID: "",
        MatchCondition: MakeMatchConditionType(),
        Direction: MakeDirectionType(),
        ActionList: MakeActionListType(),
        
    }
}



// MakeAclRuleTypeSlice() makes a slice of AclRuleType
func MakeAclRuleTypeSlice() []*AclRuleType {
    return []*AclRuleType{}
}
