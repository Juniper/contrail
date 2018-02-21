package models


// MakeLogicalInterface makes LogicalInterface
func MakeLogicalInterface() *LogicalInterface{
    return &LogicalInterface{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        LogicalInterfaceVlanTag: 0,
        LogicalInterfaceType: "",
        
    }
}

// MakeLogicalInterfaceSlice() makes a slice of LogicalInterface
func MakeLogicalInterfaceSlice() []*LogicalInterface {
    return []*LogicalInterface{}
}


