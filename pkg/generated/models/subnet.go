package models


// MakeSubnet makes Subnet
func MakeSubnet() *Subnet{
    return &Subnet{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        SubnetIPPrefix: MakeSubnetType(),
        
    }
}

// MakeSubnetSlice() makes a slice of Subnet
func MakeSubnetSlice() []*Subnet {
    return []*Subnet{}
}


