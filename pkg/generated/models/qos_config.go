package models


// MakeQosConfig makes QosConfig
func MakeQosConfig() *QosConfig{
    return &QosConfig{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        QosConfigType: "",
        MPLSExpEntries: MakeQosIdForwardingClassPairs(),
        VlanPriorityEntries: MakeQosIdForwardingClassPairs(),
        DefaultForwardingClassID: 0,
        DSCPEntries: MakeQosIdForwardingClassPairs(),
        
    }
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
    return []*QosConfig{}
}


