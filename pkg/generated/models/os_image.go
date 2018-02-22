package models


// MakeOsImage makes OsImage
func MakeOsImage() *OsImage{
    return &OsImage{
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
        Owner: "",
        ID: "",
        Size_: 0,
        Status: "",
        Location: "",
        File: "",
        Checksum: "",
        CreatedAt: "",
        UpdatedAt: "",
        ContainerFormat: "",
        DiskFormat: "",
        Protected: false,
        Visibility: "",
        Property: "",
        MinDisk: 0,
        MinRAM: 0,
        Tags: "",
        
    }
}

// MakeOsImageSlice() makes a slice of OsImage
func MakeOsImageSlice() []*OsImage {
    return []*OsImage{}
}


