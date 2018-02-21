package models


// MakeFirewallRule makes FirewallRule
func MakeFirewallRule() *FirewallRule{
    return &FirewallRule{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Endpoint1: MakeFirewallRuleEndpointType(),
        Endpoint2: MakeFirewallRuleEndpointType(),
        ActionList: MakeActionListType(),
        Service: MakeFirewallServiceType(),
        Direction: "",
        MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
        MatchTags: MakeFirewallRuleMatchTagsType(),
        
    }
}

// MakeFirewallRuleSlice() makes a slice of FirewallRule
func MakeFirewallRuleSlice() []*FirewallRule {
    return []*FirewallRule{}
}


