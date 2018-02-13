package models

// VirtualDnsRecordType

import "encoding/json"

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
