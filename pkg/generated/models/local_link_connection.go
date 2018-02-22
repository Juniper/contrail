package models


// MakeLocalLinkConnection makes LocalLinkConnection
func MakeLocalLinkConnection() *LocalLinkConnection{
    return &LocalLinkConnection{
    //TODO(nati): Apply default
    SwitchID: "",
        PortID: "",
        SwitchInfo: "",
        
    }
}

// MakeLocalLinkConnectionSlice() makes a slice of LocalLinkConnection
func MakeLocalLinkConnectionSlice() []*LocalLinkConnection {
    return []*LocalLinkConnection{}
}


