package models

// RbacRuleType

import "encoding/json"

// RbacRuleType
type RbacRuleType struct {
	RuleObject string          `json:"rule_object"`
	RulePerms  []*RbacPermType `json:"rule_perms"`
	RuleField  string          `json:"rule_field"`
}

// String returns json representation of the object
func (model *RbacRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRbacRuleType makes RbacRuleType
func MakeRbacRuleType() *RbacRuleType {
	return &RbacRuleType{
		//TODO(nati): Apply default
		RuleField:  "",
		RuleObject: "",

		RulePerms: MakeRbacPermTypeSlice(),
	}
}

// MakeRbacRuleTypeSlice() makes a slice of RbacRuleType
func MakeRbacRuleTypeSlice() []*RbacRuleType {
	return []*RbacRuleType{}
}
