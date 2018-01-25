package models

// FirewallRuleMatchTagsTypeIdList

// FirewallRuleMatchTagsTypeIdList
//proteus:generate
type FirewallRuleMatchTagsTypeIdList struct {
	TagType []int `json:"tag_type,omitempty"`
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
