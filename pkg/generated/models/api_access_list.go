package models


// MakeAPIAccessList makes APIAccessList
func MakeAPIAccessList() *APIAccessList{
    return &APIAccessList{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        APIAccessListEntries: MakeRbacRuleEntriesType(),
        
    }
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
    return []*APIAccessList{}
}


