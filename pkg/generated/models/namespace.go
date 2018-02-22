package models


// MakeNamespace makes Namespace
func MakeNamespace() *Namespace{
    return &Namespace{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        NamespaceCidr: MakeSubnetType(),
        
    }
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
    return []*Namespace{}
}


