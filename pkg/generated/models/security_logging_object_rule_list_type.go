package models

// SecurityLoggingObjectRuleListType

import "encoding/json"

type SecurityLoggingObjectRuleListType struct {
	Rule []*SecurityLoggingObjectRuleEntryType `json:"rule"`
}

func (model *SecurityLoggingObjectRuleListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType {
	return &SecurityLoggingObjectRuleListType{
		//TODO(nati): Apply default

		Rule: MakeSecurityLoggingObjectRuleEntryTypeSlice(),
	}
}

func InterfaceToSecurityLoggingObjectRuleListType(iData interface{}) *SecurityLoggingObjectRuleListType {
	data := iData.(map[string]interface{})
	return &SecurityLoggingObjectRuleListType{

		Rule: InterfaceToSecurityLoggingObjectRuleEntryTypeSlice(data["rule"]),

		//{"Title":"","Description":"List of rules along with logging rate for each rule. Both rule-uuid and rate are optional. When rule-uuid is absent then it means all rules of associated SG or network-policy","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"rule","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"rate":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Rate","GoType":"int"},"rule_uuid":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleUUID","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SecurityLoggingObjectRuleEntryType","CollectionType":"","Column":"","Item":null,"GoName":"Rule","GoType":"SecurityLoggingObjectRuleEntryType"},"GoName":"Rule","GoType":"[]*SecurityLoggingObjectRuleEntryType"}

	}
}

func InterfaceToSecurityLoggingObjectRuleListTypeSlice(data interface{}) []*SecurityLoggingObjectRuleListType {
	list := data.([]interface{})
	result := MakeSecurityLoggingObjectRuleListTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObjectRuleListType(item))
	}
	return result
}

func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
	return []*SecurityLoggingObjectRuleListType{}
}
