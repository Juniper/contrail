package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAclRuleType makes AclRuleType
// nolint
func MakeAclRuleType() *AclRuleType {
	return &AclRuleType{
		//TODO(nati): Apply default
		RuleUUID:       "",
		MatchCondition: MakeMatchConditionType(),
		Direction:      "",
		ActionList:     MakeActionListType(),
	}
}

// MakeAclRuleType makes AclRuleType
// nolint
func InterfaceToAclRuleType(i interface{}) *AclRuleType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AclRuleType{
		//TODO(nati): Apply default
		RuleUUID:       common.InterfaceToString(m["rule_uuid"]),
		MatchCondition: InterfaceToMatchConditionType(m["match_condition"]),
		Direction:      common.InterfaceToString(m["direction"]),
		ActionList:     InterfaceToActionListType(m["action_list"]),
	}
}

// MakeAclRuleTypeSlice() makes a slice of AclRuleType
// nolint
func MakeAclRuleTypeSlice() []*AclRuleType {
	return []*AclRuleType{}
}

// InterfaceToAclRuleTypeSlice() makes a slice of AclRuleType
// nolint
func InterfaceToAclRuleTypeSlice(i interface{}) []*AclRuleType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AclRuleType{}
	for _, item := range list {
		result = append(result, InterfaceToAclRuleType(item))
	}
	return result
}
