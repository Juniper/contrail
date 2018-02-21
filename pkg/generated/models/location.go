package models


// MakeLocation makes Location
func MakeLocation() *Location{
    return &Location{
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
        Type: "",
        PrivateDNSServers: "",
        PrivateNTPHosts: "",
        PrivateOspdPackageURL: "",
        PrivateOspdUserName: "",
        PrivateOspdUserPassword: "",
        PrivateOspdVMDiskGB: "",
        PrivateOspdVMName: "",
        PrivateOspdVMRAMMB: "",
        PrivateOspdVMVcpus: "",
        PrivateRedhatPoolID: "",
        PrivateRedhatSubscriptionKey: "",
        PrivateRedhatSubscriptionPasword: "",
        PrivateRedhatSubscriptionUser: "",
        GCPAccountInfo: "",
        GCPAsn: 0,
        GCPRegion: "",
        GCPSubnet: "",
        AwsAccessKey: "",
        AwsRegion: "",
        AwsSecretKey: "",
        AwsSubnet: "",
        
    }
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
    return []*Location{}
}


