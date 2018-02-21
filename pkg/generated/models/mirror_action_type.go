package models


// MakeMirrorActionType makes MirrorActionType
func MakeMirrorActionType() *MirrorActionType{
    return &MirrorActionType{
    //TODO(nati): Apply default
    NicAssistedMirroringVlan: 0,
        AnalyzerName: "",
        NHMode: "",
        JuniperHeader: false,
        UDPPort: 0,
        RoutingInstance: "",
        StaticNHHeader: MakeStaticMirrorNhType(),
        AnalyzerIPAddress: "",
        Encapsulation: "",
        AnalyzerMacAddress: "",
        NicAssistedMirroring: false,
        
    }
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
    return []*MirrorActionType{}
}


