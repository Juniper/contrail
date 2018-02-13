package models

// RbacRuleEntriesType

import "encoding/json"

// RbacRuleEntriesType
//proteus:generate
type RbacRuleEntriesType struct {
	RbacRule []*RbacRuleType `json:"rbac_rule,omitempty"`
}

// String returns json representation of the object
func (model *RbacRuleEntriesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRbacRuleEntriesType makes RbacRuleEntriesType
func MakeRbacRuleEntriesType() *RbacRuleEntriesType {
	return &RbacRuleEntriesType{
		//TODO(nati): Apply default

		RbacRule: MakeRbacRuleTypeSlice(),
	}
}

// MakeRbacRuleEntriesTypeSlice() makes a slice of RbacRuleEntriesType
func MakeRbacRuleEntriesTypeSlice() []*RbacRuleEntriesType {
	return []*RbacRuleEntriesType{}
}
