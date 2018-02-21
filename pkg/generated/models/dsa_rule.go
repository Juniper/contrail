package models


// MakeDsaRule makes DsaRule
func MakeDsaRule() *DsaRule{
    return &DsaRule{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
        
    }
}

// MakeDsaRuleSlice() makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
    return []*DsaRule{}
}


