package models


// MakeActionListType makes ActionListType
func MakeActionListType() *ActionListType{
    return &ActionListType{
    //TODO(nati): Apply default
    GatewayName: "",
        Log: false,
        Alert: false,
        QosAction: "",
        AssignRoutingInstance: "",
        MirrorTo: MakeMirrorActionType(),
        SimpleAction: "",
        ApplyService: []string{},
        
    }
}

// MakeActionListTypeSlice() makes a slice of ActionListType
func MakeActionListTypeSlice() []*ActionListType {
    return []*ActionListType{}
}


