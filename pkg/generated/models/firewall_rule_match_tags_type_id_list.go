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

// MakeFirewallRuleMatchTagsTypeIdListSlice() makes a slice of FirewallRuleMatchTagsTypeIdList
func MakeFirewallRuleMatchTagsTypeIdListSlice() []*FirewallRuleMatchTagsTypeIdList {
	return []*FirewallRuleMatchTagsTypeIdList{}
}
