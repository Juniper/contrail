package models


// MakeCustomerAttachment makes CustomerAttachment
func MakeCustomerAttachment() *CustomerAttachment{
    return &CustomerAttachment{
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

// MakeCustomerAttachmentSlice() makes a slice of CustomerAttachment
func MakeCustomerAttachmentSlice() []*CustomerAttachment {
    return []*CustomerAttachment{}
}


