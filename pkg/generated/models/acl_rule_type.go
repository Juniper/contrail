package models

// AclRuleType

// AclRuleType
//proteus:generate
type AclRuleType struct {
	RuleUUID       string              `json:"rule_uuid,omitempty"`
	MatchCondition *MatchConditionType `json:"match_condition,omitempty"`
	Direction      DirectionType       `json:"direction,omitempty"`
	ActionList     *ActionListType     `json:"action_list,omitempty"`
}

// MakeAclRuleType makes AclRuleType
func MakeAclRuleType() *AclRuleType {
	return &AclRuleType{
		//TODO(nati): Apply default
		RuleUUID:       "",
		MatchCondition: MakeMatchConditionType(),
		Direction:      MakeDirectionType(),
		ActionList:     MakeActionListType(),
	}
}

// MakeAclRuleTypeSlice() makes a slice of AclRuleType
func MakeAclRuleTypeSlice() []*AclRuleType {
	return []*AclRuleType{}
}
