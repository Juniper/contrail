package models


// MakeDomain makes Domain
func MakeDomain() *Domain{
    return &Domain{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        DomainLimits: MakeDomainLimitsType(),
        
    }
}

// MakeDomainSlice() makes a slice of Domain
func MakeDomainSlice() []*Domain {
    return []*Domain{}
}


