package models


// MakeTag makes Tag
func MakeTag() *Tag{
    return &Tag{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        TagTypeName: "",
        TagID: "",
        TagValue: "",
        
    }
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
    return []*Tag{}
}


