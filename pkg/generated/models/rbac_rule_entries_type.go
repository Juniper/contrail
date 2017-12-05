package models

// RbacRuleEntriesType

import "encoding/json"

type RbacRuleEntriesType struct {
	RbacRule []*RbacRuleType `json:"rbac_rule"`
}

func (model *RbacRuleEntriesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeRbacRuleEntriesType() *RbacRuleEntriesType {
	return &RbacRuleEntriesType{
		//TODO(nati): Apply default

		RbacRule: MakeRbacRuleTypeSlice(),
	}
}

func InterfaceToRbacRuleEntriesType(iData interface{}) *RbacRuleEntriesType {
	data := iData.(map[string]interface{})
	return &RbacRuleEntriesType{

		RbacRule: InterfaceToRbacRuleTypeSlice(data["rbac_rule"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"rule_field":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleField","GoType":"string"},"rule_object":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleObject","GoType":"string"},"rule_perms":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"role_crud":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleCrud","GoType":"string"},"role_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleName","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RbacPermType","CollectionType":"","Column":"","Item":null,"GoName":"RulePerms","GoType":"RbacPermType"},"GoName":"RulePerms","GoType":"[]*RbacPermType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RbacRuleType","CollectionType":"","Column":"","Item":null,"GoName":"RbacRule","GoType":"RbacRuleType"},"GoName":"RbacRule","GoType":"[]*RbacRuleType"}

	}
}

func InterfaceToRbacRuleEntriesTypeSlice(data interface{}) []*RbacRuleEntriesType {
	list := data.([]interface{})
	result := MakeRbacRuleEntriesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToRbacRuleEntriesType(item))
	}
	return result
}

func MakeRbacRuleEntriesTypeSlice() []*RbacRuleEntriesType {
	return []*RbacRuleEntriesType{}
}
