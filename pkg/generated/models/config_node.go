package models


// MakeConfigNode makes ConfigNode
func MakeConfigNode() *ConfigNode{
    return &ConfigNode{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ConfigNodeIPAddress: "",
        
    }
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
    return []*ConfigNode{}
}


