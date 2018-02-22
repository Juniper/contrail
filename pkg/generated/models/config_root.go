package models


// MakeConfigRoot makes ConfigRoot
func MakeConfigRoot() *ConfigRoot{
    return &ConfigRoot{
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

// MakeConfigRootSlice() makes a slice of ConfigRoot
func MakeConfigRootSlice() []*ConfigRoot {
    return []*ConfigRoot{}
}


