package models


// MakePeeringPolicy makes PeeringPolicy
func MakePeeringPolicy() *PeeringPolicy{
    return &PeeringPolicy{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        PeeringService: "",
        
    }
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
    return []*PeeringPolicy{}
}


