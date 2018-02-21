package models


// MakeStaticMirrorNhType makes StaticMirrorNhType
func MakeStaticMirrorNhType() *StaticMirrorNhType{
    return &StaticMirrorNhType{
    //TODO(nati): Apply default
    VtepDSTIPAddress: "",
        VtepDSTMacAddress: "",
        Vni: 0,
        
    }
}

// MakeStaticMirrorNhTypeSlice() makes a slice of StaticMirrorNhType
func MakeStaticMirrorNhTypeSlice() []*StaticMirrorNhType {
    return []*StaticMirrorNhType{}
}


