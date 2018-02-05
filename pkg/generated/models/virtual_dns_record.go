package models

// VirtualDNSRecord

import "encoding/json"

// VirtualDNSRecord
type VirtualDNSRecord struct {
	Perms2               *PermType2            `json:"perms2,omitempty"`
	UUID                 string                `json:"uuid,omitempty"`
	ParentUUID           string                `json:"parent_uuid,omitempty"`
	ParentType           string                `json:"parent_type,omitempty"`
	VirtualDNSRecordData *VirtualDnsRecordType `json:"virtual_DNS_record_data,omitempty"`
	Annotations          *KeyValuePairs        `json:"annotations,omitempty"`
	IDPerms              *IdPermsType          `json:"id_perms,omitempty"`
	DisplayName          string                `json:"display_name,omitempty"`
	FQName               []string              `json:"fq_name,omitempty"`
}

// String returns json representation of the object
func (model *VirtualDNSRecord) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDNSRecord makes VirtualDNSRecord
func MakeVirtualDNSRecord() *VirtualDNSRecord {
	return &VirtualDNSRecord{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		VirtualDNSRecordData: MakeVirtualDnsRecordType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		DisplayName:          "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
	return []*VirtualDNSRecord{}
}
