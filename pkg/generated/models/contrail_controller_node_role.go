package models


// MakeContrailControllerNodeRole makes ContrailControllerNodeRole
func MakeContrailControllerNodeRole() *ContrailControllerNodeRole{
    return &ContrailControllerNodeRole{
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

// MakeContrailControllerNodeRoleSlice() makes a slice of ContrailControllerNodeRole
func MakeContrailControllerNodeRoleSlice() []*ContrailControllerNodeRole {
    return []*ContrailControllerNodeRole{}
}


