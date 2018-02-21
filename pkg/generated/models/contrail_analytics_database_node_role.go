package models


// MakeContrailAnalyticsDatabaseNodeRole makes ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRole() *ContrailAnalyticsDatabaseNodeRole{
    return &ContrailAnalyticsDatabaseNodeRole{
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

// MakeContrailAnalyticsDatabaseNodeRoleSlice() makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
    return []*ContrailAnalyticsDatabaseNodeRole{}
}


