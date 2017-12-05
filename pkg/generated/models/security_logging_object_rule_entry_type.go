package models

// SecurityLoggingObjectRuleEntryType

import "encoding/json"

type SecurityLoggingObjectRuleEntryType struct {
	RuleUUID string `json:"rule_uuid"`
	Rate     int    `json:"rate"`
}

func (model *SecurityLoggingObjectRuleEntryType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeSecurityLoggingObjectRuleEntryType() *SecurityLoggingObjectRuleEntryType {
	return &SecurityLoggingObjectRuleEntryType{
		//TODO(nati): Apply default
		RuleUUID: "",
		Rate:     0,
	}
}

func InterfaceToSecurityLoggingObjectRuleEntryType(iData interface{}) *SecurityLoggingObjectRuleEntryType {
	data := iData.(map[string]interface{})
	return &SecurityLoggingObjectRuleEntryType{
		RuleUUID: data["rule_uuid"].(string),

		//{"Title":"","Description":"Rule UUID of network policy or security-group. When this is absent it implies all rules of security-group or network-policy","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RuleUUID","GoType":"string"}
		Rate: data["rate"].(int),

		//{"Title":"","Description":"Rate at which sessions are logged. When rates are specified at multiple levels, the rate which specifies highest frequency is selected","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Rate","GoType":"int"}

	}
}

func InterfaceToSecurityLoggingObjectRuleEntryTypeSlice(data interface{}) []*SecurityLoggingObjectRuleEntryType {
	list := data.([]interface{})
	result := MakeSecurityLoggingObjectRuleEntryTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObjectRuleEntryType(item))
	}
	return result
}

func MakeSecurityLoggingObjectRuleEntryTypeSlice() []*SecurityLoggingObjectRuleEntryType {
	return []*SecurityLoggingObjectRuleEntryType{}
}
