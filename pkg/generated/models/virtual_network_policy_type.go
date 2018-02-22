package models


// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyType() *VirtualNetworkPolicyType{
    return &VirtualNetworkPolicyType{
    //TODO(nati): Apply default
    Timer: MakeTimerType(),
        Sequence: MakeSequenceType(),
        
    }
}

// MakeVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
    return []*VirtualNetworkPolicyType{}
}


