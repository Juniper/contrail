package models

// FirewallRuleMatchTagsTypeIdList

import "encoding/json"

// FirewallRuleMatchTagsTypeIdList
type FirewallRuleMatchTagsTypeIdList struct {
	TagType []int `json:"tag_type"`
}

// String returns json representation of the object
func (model *FirewallRuleMatchTagsTypeIdList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallRuleMatchTagsTypeIdList makes FirewallRuleMatchTagsTypeIdList
func MakeFirewallRuleMatchTagsTypeIdList() *FirewallRuleMatchTagsTypeIdList {
	return &FirewallRuleMatchTagsTypeIdList{
		//TODO(nati): Apply default

		TagType: []int{},
	}
}

// InterfaceToFirewallRuleMatchTagsTypeIdList makes FirewallRuleMatchTagsTypeIdList from interface
func InterfaceToFirewallRuleMatchTagsTypeIdList(iData interface{}) *FirewallRuleMatchTagsTypeIdList {
	data := iData.(map[string]interface{})
	return &FirewallRuleMatchTagsTypeIdList{

		TagType: data["tag_type"].([]int),

		//{"type":"array","item":{"type":"integer"}}

	}
}

// InterfaceToFirewallRuleMatchTagsTypeIdListSlice makes a slice of FirewallRuleMatchTagsTypeIdList from interface
func InterfaceToFirewallRuleMatchTagsTypeIdListSlice(data interface{}) []*FirewallRuleMatchTagsTypeIdList {
	list := data.([]interface{})
	result := MakeFirewallRuleMatchTagsTypeIdListSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRuleMatchTagsTypeIdList(item))
	}
	return result
}

// MakeFirewallRuleMatchTagsTypeIdListSlice() makes a slice of FirewallRuleMatchTagsTypeIdList
func MakeFirewallRuleMatchTagsTypeIdListSlice() []*FirewallRuleMatchTagsTypeIdList {
	return []*FirewallRuleMatchTagsTypeIdList{}
}
