package models


// MakeNetworkDeviceConfig makes NetworkDeviceConfig
func MakeNetworkDeviceConfig() *NetworkDeviceConfig{
    return &NetworkDeviceConfig{
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

// MakeNetworkDeviceConfigSlice() makes a slice of NetworkDeviceConfig
func MakeNetworkDeviceConfigSlice() []*NetworkDeviceConfig {
    return []*NetworkDeviceConfig{}
}


