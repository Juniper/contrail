package models


// MakeBridgeDomain makes BridgeDomain
func MakeBridgeDomain() *BridgeDomain{
    return &BridgeDomain{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        MacAgingTime: 0,
        Isid: 0,
        MacLearningEnabled: false,
        MacMoveControl: MakeMACMoveLimitControlType(),
        MacLimitControl: MakeMACLimitControlType(),
        
    }
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
    return []*BridgeDomain{}
}


