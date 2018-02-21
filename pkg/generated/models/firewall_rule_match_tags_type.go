package models


// MakeFirewallRuleMatchTagsType makes FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsType() *FirewallRuleMatchTagsType{
    return &FirewallRuleMatchTagsType{
    //TODO(nati): Apply default
    TagList: []string{},
        
    }
}

// MakeFirewallRuleMatchTagsTypeSlice() makes a slice of FirewallRuleMatchTagsType
func MakeFirewallRuleMatchTagsTypeSlice() []*FirewallRuleMatchTagsType {
    return []*FirewallRuleMatchTagsType{}
}


