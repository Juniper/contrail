package models

// RbacRuleType

// RbacRuleType
//proteus:generate
type RbacRuleType struct {
	RuleObject string          `json:"rule_object,omitempty"`
	RulePerms  []*RbacPermType `json:"rule_perms,omitempty"`
	RuleField  string          `json:"rule_field,omitempty"`
}

// MakeRbacRuleType makes RbacRuleType
func MakeRbacRuleType() *RbacRuleType {
	return &RbacRuleType{
		//TODO(nati): Apply default
		RuleObject: "",

		RulePerms: MakeRbacPermTypeSlice(),

		RuleField: "",
	}
}

// MakeRbacRuleTypeSlice() makes a slice of RbacRuleType
func MakeRbacRuleTypeSlice() []*RbacRuleType {
	return []*RbacRuleType{}
}
