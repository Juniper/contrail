package models

// VirtualDnsRecordType

import "encoding/json"

// VirtualDnsRecordType
type VirtualDnsRecordType struct {
	RecordName         string             `json:"record_name"`
	RecordClass        DnsRecordClassType `json:"record_class"`
	RecordData         string             `json:"record_data"`
	RecordType         DnsRecordTypeType  `json:"record_type"`
	RecordTTLSeconds   int                `json:"record_ttl_seconds"`
	RecordMXPreference int                `json:"record_mx_preference"`
}

// String returns json representation of the object
func (model *VirtualDnsRecordType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
