package models


// MakeServer makes Server
func MakeServer() *Server{
    return &Server{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Created: "",
        HostId: "",
        ID: "",
        Name: "",
        Image: MakeOpenStackImageProperty(),
        Flavor: MakeOpenStackFlavorProperty(),
        Addresses: MakeOpenStackAddress(),
        AccessIPv4: "",
        AccessIPv6: "",
        ConfigDrive: false,
        Progress: 0,
        Status: "",
        HostStatus: "",
        TenantID: "",
        Updated: "",
        UserID: 0,
        Locked: false,
        
    }
}

// MakeServerSlice() makes a slice of Server
func MakeServerSlice() []*Server {
    return []*Server{}
}


