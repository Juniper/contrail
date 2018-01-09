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
		RuleObject: "",

		RulePerms: MakeRbacPermTypeSlice(),

		RuleField: "",
	}
}

// InterfaceToRbacRuleType makes RbacRuleType from interface
func InterfaceToRbacRuleType(iData interface{}) *RbacRuleType {
	data := iData.(map[string]interface{})
	return &RbacRuleType{
		RuleField: data["rule_field"].(string),

		//{"description":"Name of the level one field (property) for this object, * represent all properties","type":"string"}
		RuleObject: data["rule_object"].(string),

		//{"description":"Name of the REST API (object) for this rule, * represent all objects","type":"string"}

		RulePerms: InterfaceToRbacPermTypeSlice(data["rule_perms"]),

		//{"description":"List of [(role, permissions),...]","type":"array","item":{"type":"object","properties":{"role_crud":{"type":"string"},"role_name":{"type":"string"}}}}

	}
}

// InterfaceToRbacRuleTypeSlice makes a slice of RbacRuleType from interface
func InterfaceToRbacRuleTypeSlice(data interface{}) []*RbacRuleType {
	list := data.([]interface{})
	result := MakeRbacRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleType(item))
	}
	return result
}

// MakeRbacRuleTypeSlice() makes a slice of RbacRuleType
func MakeRbacRuleTypeSlice() []*RbacRuleType {
	return []*RbacRuleType{}
}
