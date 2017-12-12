package models

// SecurityLoggingObjectRuleListType

import "encoding/json"

// SecurityLoggingObjectRuleListType
type SecurityLoggingObjectRuleListType struct {
	Rule []*SecurityLoggingObjectRuleEntryType `json:"rule"`
}

// String returns json representation of the object
func (model *SecurityLoggingObjectRuleListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListType() *SecurityLoggingObjectRuleListType {
	return &SecurityLoggingObjectRuleListType{
		//TODO(nati): Apply default

		Rule: MakeSecurityLoggingObjectRuleEntryTypeSlice(),
	}
}

// InterfaceToSecurityLoggingObjectRuleListType makes SecurityLoggingObjectRuleListType from interface
func InterfaceToSecurityLoggingObjectRuleListType(iData interface{}) *SecurityLoggingObjectRuleListType {
	data := iData.(map[string]interface{})
	return &SecurityLoggingObjectRuleListType{

		Rule: InterfaceToSecurityLoggingObjectRuleEntryTypeSlice(data["rule"]),

		//{"description":"List of rules along with logging rate for each rule. Both rule-uuid and rate are optional. When rule-uuid is absent then it means all rules of associated SG or network-policy","type":"array","item":{"type":"object","properties":{"rate":{"type":"integer"},"rule_uuid":{"type":"string"}}}}

	}
}

// InterfaceToSecurityLoggingObjectRuleListTypeSlice makes a slice of SecurityLoggingObjectRuleListType from interface
func InterfaceToSecurityLoggingObjectRuleListTypeSlice(data interface{}) []*SecurityLoggingObjectRuleListType {
	list := data.([]interface{})
	result := MakeSecurityLoggingObjectRuleListTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityLoggingObjectRuleListType(item))
	}
	return result
}

// MakeSecurityLoggingObjectRuleListTypeSlice() makes a slice of SecurityLoggingObjectRuleListType
func MakeSecurityLoggingObjectRuleListTypeSlice() []*SecurityLoggingObjectRuleListType {
	return []*SecurityLoggingObjectRuleListType{}
}
