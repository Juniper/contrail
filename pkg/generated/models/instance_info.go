package models


// MakeInstanceInfo makes InstanceInfo
func MakeInstanceInfo() *InstanceInfo{
    return &InstanceInfo{
    //TODO(nati): Apply default
    DisplayName: "",
        ImageSource: "",
        LocalGB: "",
        MemoryMB: "",
        NovaHostID: "",
        RootGB: "",
        SwapMB: "",
        Vcpus: "",
        Capabilities: "",
        
    }
}

// MakeInstanceInfoSlice() makes a slice of InstanceInfo
func MakeInstanceInfoSlice() []*InstanceInfo {
    return []*InstanceInfo{}
}


