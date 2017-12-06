package models

// RbacRuleEntriesType

import "encoding/json"

// RbacRuleEntriesType
type RbacRuleEntriesType struct {
	RbacRule []*RbacRuleType `json:"rbac_rule"`
}

//  parents relation object

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

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"rule_field":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleField","GoType":"string","GoPremitive":true},"rule_object":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleObject","GoType":"string","GoPremitive":true},"rule_perms":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"role_crud":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleCrud","GoType":"string","GoPremitive":true},"role_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoleName","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RbacPermType","CollectionType":"","Column":"","Item":null,"GoName":"RulePerms","GoType":"RbacPermType","GoPremitive":false},"GoName":"RulePerms","GoType":"[]*RbacPermType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RbacRuleType","CollectionType":"","Column":"","Item":null,"GoName":"RbacRule","GoType":"RbacRuleType","GoPremitive":false},"GoName":"RbacRule","GoType":"[]*RbacRuleType","GoPremitive":true}

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
