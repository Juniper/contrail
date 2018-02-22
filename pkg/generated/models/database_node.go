package models


// MakeDatabaseNode makes DatabaseNode
func MakeDatabaseNode() *DatabaseNode{
    return &DatabaseNode{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        DatabaseNodeIPAddress: "",
        
    }
}

// MakeDatabaseNodeSlice() makes a slice of DatabaseNode
func MakeDatabaseNodeSlice() []*DatabaseNode {
    return []*DatabaseNode{}
}


