package models


// MakePolicyManagement makes PolicyManagement
func MakePolicyManagement() *PolicyManagement{
    return &PolicyManagement{
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

// MakePolicyManagementSlice() makes a slice of PolicyManagement
func MakePolicyManagementSlice() []*PolicyManagement {
    return []*PolicyManagement{}
}


