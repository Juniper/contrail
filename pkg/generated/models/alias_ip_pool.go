package models


// MakeAliasIPPool makes AliasIPPool
func MakeAliasIPPool() *AliasIPPool{
    return &AliasIPPool{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
    }
}

// MakeAliasIPPoolSlice() makes a slice of AliasIPPool
func MakeAliasIPPoolSlice() []*AliasIPPool {
    return []*AliasIPPool{}
}


