package models

// RbacRuleEntriesType

import "encoding/json"

// RbacRuleEntriesType
type RbacRuleEntriesType struct {
	RbacRule []*RbacRuleType `json:"rbac_rule"`
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

// InterfaceToRbacRuleEntriesType makes RbacRuleEntriesType from interface
func InterfaceToRbacRuleEntriesType(iData interface{}) *RbacRuleEntriesType {
	data := iData.(map[string]interface{})
	return &RbacRuleEntriesType{

		RbacRule: InterfaceToRbacRuleTypeSlice(data["rbac_rule"]),

		//{"type":"array","item":{"type":"object","properties":{"rule_field":{"type":"string"},"rule_object":{"type":"string"},"rule_perms":{"type":"array","item":{"type":"object","properties":{"role_crud":{"type":"string"},"role_name":{"type":"string"}}}}}}}

	}
}

// InterfaceToRbacRuleEntriesTypeSlice makes a slice of RbacRuleEntriesType from interface
func InterfaceToRbacRuleEntriesTypeSlice(data interface{}) []*RbacRuleEntriesType {
	list := data.([]interface{})
	result := MakeRbacRuleEntriesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleEntriesType(item))
	}
	return result
}

// MakeRbacRuleEntriesTypeSlice() makes a slice of RbacRuleEntriesType
func MakeRbacRuleEntriesTypeSlice() []*RbacRuleEntriesType {
	return []*RbacRuleEntriesType{}
}
