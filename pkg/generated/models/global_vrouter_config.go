package models


// MakeGlobalVrouterConfig makes GlobalVrouterConfig
func MakeGlobalVrouterConfig() *GlobalVrouterConfig{
    return &GlobalVrouterConfig{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
        FlowAgingTimeoutList: MakeFlowAgingTimeoutList(),
        ForwardingMode: "",
        FlowExportRate: 0,
        LinklocalServices: MakeLinklocalServicesTypes(),
        EncapsulationPriorities: MakeEncapsulationPrioritiesType(),
        VxlanNetworkIdentifierMode: "",
        EnableSecurityLogging: false,
        
    }
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
    return []*GlobalVrouterConfig{}
}


