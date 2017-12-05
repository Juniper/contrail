package models

// FirewallRuleMatchTagsType

import "encoding/json"

type FirewallRuleMatchTagsType struct {
	TagList []string `json:"tag_list"`
}

func (model *FirewallRuleMatchTagsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFirewallRuleMatchTagsType() *FirewallRuleMatchTagsType {
	return &FirewallRuleMatchTagsType{
		//TODO(nati): Apply default
		TagList: []string{},
	}
}

func InterfaceToFirewallRuleMatchTagsType(iData interface{}) *FirewallRuleMatchTagsType {
	data := iData.(map[string]interface{})
	return &FirewallRuleMatchTagsType{
		TagList: data["tag_list"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagList","GoType":"string"},"GoName":"TagList","GoType":"[]string"}

	}
}

func InterfaceToFirewallRuleMatchTagsTypeSlice(data interface{}) []*FirewallRuleMatchTagsType {
	list := data.([]interface{})
	result := MakeFirewallRuleMatchTagsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleMatchTagsType(item))
	}
	return result
}

func MakeFirewallRuleMatchTagsTypeSlice() []*FirewallRuleMatchTagsType {
	return []*FirewallRuleMatchTagsType{}
}
