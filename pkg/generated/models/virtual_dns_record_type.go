package models

// VirtualDnsRecordType

import "encoding/json"

// VirtualDnsRecordType
type VirtualDnsRecordType struct {
	RecordMXPreference int                `json:"record_mx_preference,omitempty"`
	RecordName         string             `json:"record_name,omitempty"`
	RecordClass        DnsRecordClassType `json:"record_class,omitempty"`
	RecordData         string             `json:"record_data,omitempty"`
	RecordType         DnsRecordTypeType  `json:"record_type,omitempty"`
	RecordTTLSeconds   int                `json:"record_ttl_seconds,omitempty"`
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
		RecordTTLSeconds:   0,
		RecordMXPreference: 0,
		RecordName:         "",
		RecordClass:        MakeDnsRecordClassType(),
		RecordData:         "",
		RecordType:         MakeDnsRecordTypeType(),
	}
}

// MakeVirtualDnsRecordTypeSlice() makes a slice of VirtualDnsRecordType
func MakeVirtualDnsRecordTypeSlice() []*VirtualDnsRecordType {
	return []*VirtualDnsRecordType{}
}
