package models

// FirewallRuleMatchTagsType

import "encoding/json"

// FirewallRuleMatchTagsType
//proteus:generate
type FirewallRuleMatchTagsType struct {
	TagList []string `json:"tag_list,omitempty"`
}

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

// MakeFirewallRuleMatchTagsTypeSlice() makes a slice of FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsTypeSlice() []*FirewallRuleMatchTagsType {
	return []*FirewallRuleMatchTagsType{}
}
