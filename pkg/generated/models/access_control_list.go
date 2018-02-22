package models


// MakeAccessControlList makes AccessControlList
func MakeAccessControlList() *AccessControlList{
    return &AccessControlList{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        AccessControlListHash: 0,
        AccessControlListEntries: MakeAclEntriesType(),
        
    }
}

// MakeAccessControlListSlice() makes a slice of AccessControlList
func MakeAccessControlListSlice() []*AccessControlList {
    return []*AccessControlList{}
}


