package models

// VirtualDnsRecordType

// VirtualDnsRecordType
//proteus:generate
type VirtualDnsRecordType struct {
	RecordName         string             `json:"record_name,omitempty"`
	RecordClass        DnsRecordClassType `json:"record_class,omitempty"`
	RecordData         string             `json:"record_data,omitempty"`
	RecordType         DnsRecordTypeType  `json:"record_type,omitempty"`
	RecordTTLSeconds   int                `json:"record_ttl_seconds,omitempty"`
	RecordMXPreference int                `json:"record_mx_preference,omitempty"`
}

// MakeVirtualDnsRecordType makes VirtualDnsRecordType
func MakeVirtualDnsRecordType() *VirtualDnsRecordType {
	return &VirtualDnsRecordType{
		//TODO(nati): Apply default
		RecordName:         "",
		RecordClass:        MakeDnsRecordClassType(),
		RecordData:         "",
		RecordType:         MakeDnsRecordTypeType(),
		RecordTTLSeconds:   0,
		RecordMXPreference: 0,
	}
}

// MakeVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
func MakeVirtualDnsRecordTypeSlice() []*VirtualDnsRecordType {
	return []*VirtualDnsRecordType{}
}
