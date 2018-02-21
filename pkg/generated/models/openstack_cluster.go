package models


// MakeOpenstackCluster makes OpenstackCluster
func MakeOpenstackCluster() *OpenstackCluster{
    return &OpenstackCluster{
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
        AdminPassword: "",
        ContrailClusterID: "",
        DefaultCapacityDrives: "",
        DefaultJournalDrives: "",
        DefaultOsdDrives: "",
        DefaultPerformanceDrives: "",
        DefaultStorageAccessBondInterfaceMembers: "",
        DefaultStorageBackendBondInterfaceMembers: "",
        ExternalAllocationPoolEnd: "",
        ExternalAllocationPoolStart: "",
        ExternalNetCidr: "",
        OpenstackWebui: "",
        PublicGateway: "",
        PublicIP: "",
        
    }
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
    return []*OpenstackCluster{}
}


