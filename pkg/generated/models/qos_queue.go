package models


// MakeQosQueue makes QosQueue
func MakeQosQueue() *QosQueue{
    return &QosQueue{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        QosQueueIdentifier: 0,
        MaxBandwidth: 0,
        MinBandwidth: 0,
        
    }
}

// MakeQosQueueSlice() makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
    return []*QosQueue{}
}


