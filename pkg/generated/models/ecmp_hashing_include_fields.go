package models


// MakeEcmpHashingIncludeFields makes EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFields() *EcmpHashingIncludeFields{
    return &EcmpHashingIncludeFields{
    //TODO(nati): Apply default
    DestinationIP: false,
        IPProtocol: false,
        SourceIP: false,
        HashingConfigured: false,
        SourcePort: false,
        DestinationPort: false,
        
    }
}

// MakeEcmpHashingIncludeFieldsSlice() makes a slice of EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
    return []*EcmpHashingIncludeFields{}
}


