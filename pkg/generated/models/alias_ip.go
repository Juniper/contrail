package models


// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP{
    return &AliasIP{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        AliasIPAddress: "",
        AliasIPAddressFamily: "",
        
    }
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
    return []*AliasIP{}
}


