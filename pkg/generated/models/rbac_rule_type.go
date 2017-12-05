package models

// RbacRuleType

import "encoding/json"

type RbacRuleType struct {
	RuleObject string          `json:"rule_object"`
	RulePerms  []*RbacPermType `json:"rule_perms"`
	RuleField  string          `json:"rule_field"`
}

func (model *RbacRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeRbacRuleType() *RbacRuleType {
	return &RbacRuleType{
		//TODO(nati): Apply default
		RuleObject: "",

		RulePerms: MakeRbacPermTypeSlice(),

		RuleField: "",
	}
}

func InterfaceToRbacRuleType(iData interface{}) *RbacRuleType {
	data := iData.(map[string]interface{})
	return &RbacRuleType{
		RuleObject: data["rule_object"].(string),

		//{"Title":"","Description":"Name of the REST API (object) for this rule, * represent all objects","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleObject","GoType":"string"}

		RulePerms: InterfaceToRbacPermTypeSlice(data["rule_perms"]),

		//{"Title":"","Description":"List of [(role, permissions),...]","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"role_crud":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleCrud","GoType":"string"},"role_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleName","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RbacPermType","CollectionType":"","Column":"","Item":null,"GoName":"RulePerms","GoType":"RbacPermType"},"GoName":"RulePerms","GoType":"[]*RbacPermType"}
		RuleField: data["rule_field"].(string),

		//{"Title":"","Description":"Name of the level one field (property) for this object, * represent all properties","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleField","GoType":"string"}

	}
}

func InterfaceToRbacRuleTypeSlice(data interface{}) []*RbacRuleType {
	list := data.([]interface{})
	result := MakeRbacRuleTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleType(item))
	}
	return result
}

func MakeRbacRuleTypeSlice() []*RbacRuleType {
	return []*RbacRuleType{}
}
