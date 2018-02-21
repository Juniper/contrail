package models


// MakeAddressGroup makes AddressGroup
func MakeAddressGroup() *AddressGroup{
    return &AddressGroup{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        AddressGroupPrefix: MakeSubnetListType(),
        
    }
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
    return []*AddressGroup{}
}


