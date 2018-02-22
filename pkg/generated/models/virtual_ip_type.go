package models


// MakeVirtualIpType makes VirtualIpType
func MakeVirtualIpType() *VirtualIpType{
    return &VirtualIpType{
    //TODO(nati): Apply default
    Status: "",
        StatusDescription: "",
        Protocol: "",
        SubnetID: "",
        PersistenceCookieName: "",
        ConnectionLimit: 0,
        PersistenceType: "",
        AdminState: false,
        Address: "",
        ProtocolPort: 0,
        
    }
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
    return []*VirtualIpType{}
}


