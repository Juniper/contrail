package models

// FirewallRuleMatchTagsTypeIdList

import "encoding/json"

type FirewallRuleMatchTagsTypeIdList struct {
	TagType []int `json:"tag_type"`
}

func (model *FirewallRuleMatchTagsTypeIdList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFirewallRuleMatchTagsTypeIdList() *FirewallRuleMatchTagsTypeIdList {
	return &FirewallRuleMatchTagsTypeIdList{
		//TODO(nati): Apply default

		TagType: []int{},
	}
}

func InterfaceToFirewallRuleMatchTagsTypeIdList(iData interface{}) *FirewallRuleMatchTagsTypeIdList {
	data := iData.(map[string]interface{})
	return &FirewallRuleMatchTagsTypeIdList{

		TagType: data["tag_type"].([]int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagType","GoType":"int"},"GoName":"TagType","GoType":"[]int"}

	}
}

func InterfaceToFirewallRuleMatchTagsTypeIdListSlice(data interface{}) []*FirewallRuleMatchTagsTypeIdList {
	list := data.([]interface{})
	result := MakeFirewallRuleMatchTagsTypeIdListSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleMatchTagsTypeIdList(item))
	}
	return result
}

func MakeFirewallRuleMatchTagsTypeIdListSlice() []*FirewallRuleMatchTagsTypeIdList {
	return []*FirewallRuleMatchTagsTypeIdList{}
}
