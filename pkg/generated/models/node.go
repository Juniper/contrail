package models


// MakeNode makes Node
func MakeNode() *Node{
    return &Node{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Hostname: "",
        IPAddress: "",
        MacAddress: "",
        Type: "",
        Password: "",
        SSHKey: "",
        Username: "",
        AwsAmi: "",
        AwsInstanceType: "",
        GCPImage: "",
        GCPMachineType: "",
        PrivateMachineProperties: "",
        PrivateMachineState: "",
        PrivatePowerManagementIP: "",
        PrivatePowerManagementPassword: "",
        PrivatePowerManagementUsername: "",
        
    }
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
    return []*Node{}
}


