package models


// MakeVirtualDnsRecordType makes VirtualDnsRecordType
func MakeVirtualDnsRecordType() *VirtualDnsRecordType{
    return &VirtualDnsRecordType{
    //TODO(nati): Apply default
    RecordName: "",
        RecordClass: "",
        RecordData: "",
        RecordType: "",
        RecordTTLSeconds: 0,
        RecordMXPreference: 0,
        
    }
}

// MakeVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
func MakeVirtualDnsRecordTypeSlice() []*VirtualDnsRecordType {
    return []*VirtualDnsRecordType{}
}


