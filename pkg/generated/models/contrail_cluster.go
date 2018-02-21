package models


// MakeContrailCluster makes ContrailCluster
func MakeContrailCluster() *ContrailCluster{
    return &ContrailCluster{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ConfigAuditTTL: "",
        ContrailWebui: "",
        DataTTL: "",
        DefaultGateway: "",
        DefaultVrouterBondInterface: "",
        DefaultVrouterBondInterfaceMembers: "",
        FlowTTL: "",
        StatisticsTTL: "",
        
    }
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
    return []*ContrailCluster{}
}


