package models


// MakeSecurityGroup makes SecurityGroup
func MakeSecurityGroup() *SecurityGroup{
    return &SecurityGroup{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        SecurityGroupEntries: MakePolicyEntriesType(),
        ConfiguredSecurityGroupID: 0,
        SecurityGroupID: 0,
        
    }
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
    return []*SecurityGroup{}
}


