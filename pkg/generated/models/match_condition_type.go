package models


// MakeMatchConditionType makes MatchConditionType
func MakeMatchConditionType() *MatchConditionType{
    return &MatchConditionType{
    //TODO(nati): Apply default
    SRCPort: MakePortType(),
        SRCAddress: MakeAddressType(),
        Ethertype: "",
        DSTAddress: MakeAddressType(),
        DSTPort: MakePortType(),
        Protocol: "",
        
    }
}

// MakeMatchConditionTypeSlice() makes a slice of MatchConditionType
func MakeMatchConditionTypeSlice() []*MatchConditionType {
    return []*MatchConditionType{}
}


