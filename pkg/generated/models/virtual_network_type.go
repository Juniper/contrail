package models


// MakeVirtualNetworkType makes VirtualNetworkType
func MakeVirtualNetworkType() *VirtualNetworkType{
    return &VirtualNetworkType{
    //TODO(nati): Apply default
    ForwardingMode: "",
        AllowTransit: false,
        NetworkID: 0,
        MirrorDestination: false,
        VxlanNetworkIdentifier: 0,
        RPF: "",
        
    }
}

// MakeVirtualNetworkTypeSlice() makes a slice of VirtualNetworkType
func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
    return []*VirtualNetworkType{}
}


