package models


// MakeFlavor makes Flavor
func MakeFlavor() *Flavor{
    return &Flavor{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Name: "",
        Disk: 0,
        Vcpus: 0,
        RAM: 0,
        ID: "",
        Property: "",
        RXTXFactor: 0,
        Swap: 0,
        IsPublic: false,
        Ephemeral: 0,
        Links: MakeOpenStackLink(),
        
    }
}

// MakeFlavorSlice() makes a slice of Flavor
func MakeFlavorSlice() []*Flavor {
    return []*Flavor{}
}


