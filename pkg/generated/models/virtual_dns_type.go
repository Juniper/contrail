package models


// MakeVirtualDnsType makes VirtualDnsType
func MakeVirtualDnsType() *VirtualDnsType{
    return &VirtualDnsType{
    //TODO(nati): Apply default
    FloatingIPRecord: "",
        DomainName: "",
        ExternalVisible: false,
        NextVirtualDNS: "",
        DynamicRecordsFromClient: false,
        ReverseResolution: false,
        DefaultTTLSeconds: 0,
        RecordOrder: "",
        
    }
}

// MakeVirtualDnsTypeSlice() makes a slice of VirtualDnsType
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
    return []*VirtualDnsType{}
}


