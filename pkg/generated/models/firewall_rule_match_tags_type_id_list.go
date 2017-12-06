package models

// FirewallRuleMatchTagsTypeIdList

import "encoding/json"

// FirewallRuleMatchTagsTypeIdList
type FirewallRuleMatchTagsTypeIdList struct {
	TagType []int `json:"tag_type"`
}

//  parents relation object

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

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"TagType","GoType":"int","GoPremitive":true},"GoName":"TagType","GoType":"[]int","GoPremitive":true}

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
