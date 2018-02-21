package models


// MakeControllerNodeRole makes ControllerNodeRole
func MakeControllerNodeRole() *ControllerNodeRole{
    return &ControllerNodeRole{
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
        CapacityDrives: "",
        InternalapiBondInterfaceMembers: "",
        PerformanceDrives: "",
        StorageManagementBondInterfaceMembers: "",
        
    }
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
    return []*ControllerNodeRole{}
}


