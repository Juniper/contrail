package models

// SecurityLoggingObjectRuleEntryType

import "encoding/json"

// SecurityLoggingObjectRuleEntryType
type SecurityLoggingObjectRuleEntryType struct {
	RuleUUID string `json:"rule_uuid"`
	Rate     int    `json:"rate"`
}

// String returns json representation of the object
func (model *SecurityLoggingObjectRuleEntryType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryType() *SecurityLoggingObjectRuleEntryType {
	return &SecurityLoggingObjectRuleEntryType{
		//TODO(nati): Apply default
		Rate:     0,
		RuleUUID: "",
	}
}

// InterfaceToSecurityLoggingObjectRuleEntryType makes SecurityLoggingObjectRuleEntryType from interface
func InterfaceToSecurityLoggingObjectRuleEntryType(iData interface{}) *SecurityLoggingObjectRuleEntryType {
	data := iData.(map[string]interface{})
	return &SecurityLoggingObjectRuleEntryType{
		RuleUUID: data["rule_uuid"].(string),

		//{"description":"Rule UUID of network policy or security-group. When this is absent it implies all rules of security-group or network-policy","type":"string"}
		Rate: data["rate"].(int),

		//{"description":"Rate at which sessions are logged. When rates are specified at multiple levels, the rate which specifies highest frequency is selected","type":"integer"}

	}
}

// InterfaceToSecurityLoggingObjectRuleEntryTypeSlice makes a slice of SecurityLoggingObjectRuleEntryType from interface
func InterfaceToSecurityLoggingObjectRuleEntryTypeSlice(data interface{}) []*SecurityLoggingObjectRuleEntryType {
	list := data.([]interface{})
	result := MakeSecurityLoggingObjectRuleEntryTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObjectRuleEntryType(item))
	}
	return result
}

// MakeSecurityLoggingObjectRuleEntryTypeSlice() makes a slice of SecurityLoggingObjectRuleEntryType
func MakeSecurityLoggingObjectRuleEntryTypeSlice() []*SecurityLoggingObjectRuleEntryType {
	return []*SecurityLoggingObjectRuleEntryType{}
}
