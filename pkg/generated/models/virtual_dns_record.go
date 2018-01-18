package models

// VirtualDNSRecord

import "encoding/json"

// VirtualDNSRecord
type VirtualDNSRecord struct {
	VirtualDNSRecordData *VirtualDnsRecordType `json:"virtual_DNS_record_data,omitempty"`
	IDPerms              *IdPermsType          `json:"id_perms,omitempty"`
	DisplayName          string                `json:"display_name,omitempty"`
	UUID                 string                `json:"uuid,omitempty"`
	ParentUUID           string                `json:"parent_uuid,omitempty"`
	ParentType           string                `json:"parent_type,omitempty"`
	Annotations          *KeyValuePairs        `json:"annotations,omitempty"`
	Perms2               *PermType2            `json:"perms2,omitempty"`
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
		DisplayName:          "",
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		VirtualDNSRecordData: MakeVirtualDnsRecordType(),
		IDPerms:              MakeIdPermsType(),
		FQName:               []string{},
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
	}
}

// MakeVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
	return []*VirtualDNSRecord{}
}
