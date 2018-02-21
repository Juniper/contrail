package models


// MakeAppformixNodeRole makes AppformixNodeRole
func MakeAppformixNodeRole() *AppformixNodeRole{
    return &AppformixNodeRole{
    //TODO(nati): Apply default
    ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ProvisioningStartTime: "",
        ProvisioningState: "",
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

// MakeAppformixNodeRoleSlice() makes a slice of AppformixNodeRole
func MakeAppformixNodeRoleSlice() []*AppformixNodeRole {
    return []*AppformixNodeRole{}
}


