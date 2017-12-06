package models

// FirewallRuleMatchTagsType

import "encoding/json"

// FirewallRuleMatchTagsType
type FirewallRuleMatchTagsType struct {
	TagList []string `json:"tag_list"`
}

//  parents relation object

// String returns json representation of the object
func (model *FirewallRuleMatchTagsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsType() *FirewallRuleMatchTagsType {
	return &FirewallRuleMatchTagsType{
		//TODO(nati): Apply default
		TagList: []string{},
	}
}

// InterfaceToFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType from interface
func InterfaceToFirewallRuleMatchTagsType(iData interface{}) *FirewallRuleMatchTagsType {
	data := iData.(map[string]interface{})
	return &FirewallRuleMatchTagsType{
		TagList: data["tag_list"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagList","GoType":"string","GoPremitive":true},"GoName":"TagList","GoType":"[]string","GoPremitive":true}

	}
}

// InterfaceToFirewallRuleMatchTagsTypeSlice makes a slice of FirewallRuleMatchTagsType from interface
func InterfaceToFirewallRuleMatchTagsTypeSlice(data interface{}) []*FirewallRuleMatchTagsType {
	list := data.([]interface{})
	result := MakeFirewallRuleMatchTagsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleMatchTagsType(item))
	}
	return result
}

// MakeFirewallRuleMatchTagsTypeSlice() makes a slice of FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsTypeSlice() []*FirewallRuleMatchTagsType {
	return []*FirewallRuleMatchTagsType{}
}
