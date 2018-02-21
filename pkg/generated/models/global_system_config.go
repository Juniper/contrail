package models


// MakeGlobalSystemConfig makes GlobalSystemConfig
func MakeGlobalSystemConfig() *GlobalSystemConfig{
    return &GlobalSystemConfig{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ConfigVersion: "",
        BgpaasParameters: MakeBGPaaServiceParametersType(),
        AlarmEnable: false,
        MacMoveControl: MakeMACMoveLimitControlType(),
        PluginTuning: MakePluginProperties(),
        IbgpAutoMesh: false,
        MacAgingTime: 0,
        BGPAlwaysCompareMed: false,
        UserDefinedLogStatistics: MakeUserDefinedLogStatList(),
        GracefulRestartParameters: MakeGracefulRestartParametersType(),
        IPFabricSubnets: MakeSubnetListType(),
        AutonomousSystem: 0,
        MacLimitControl: MakeMACLimitControlType(),
        
    }
}

// MakeGlobalSystemConfigSlice() makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
    return []*GlobalSystemConfig{}
}


