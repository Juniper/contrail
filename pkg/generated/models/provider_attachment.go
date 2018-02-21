package models


// MakeProviderAttachment makes ProviderAttachment
func MakeProviderAttachment() *ProviderAttachment{
    return &ProviderAttachment{
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

// MakeProviderAttachmentSlice() makes a slice of ProviderAttachment
func MakeProviderAttachmentSlice() []*ProviderAttachment {
    return []*ProviderAttachment{}
}


