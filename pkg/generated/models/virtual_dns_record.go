package models

// VirtualDNSRecord

import "encoding/json"

// VirtualDNSRecord
type VirtualDNSRecord struct {
	Annotations          *KeyValuePairs        `json:"annotations,omitempty"`
	Perms2               *PermType2            `json:"perms2,omitempty"`
	ParentType           string                `json:"parent_type,omitempty"`
	ParentUUID           string                `json:"parent_uuid,omitempty"`
	FQName               []string              `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType          `json:"id_perms,omitempty"`
	VirtualDNSRecordData *VirtualDnsRecordType `json:"virtual_DNS_record_data,omitempty"`
	DisplayName          string                `json:"display_name,omitempty"`
	UUID                 string                `json:"uuid,omitempty"`
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
		VirtualDNSRecordData: MakeVirtualDnsRecordType(),
		DisplayName:          "",
		UUID:                 "",
		ParentUUID:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ParentType:           "",
	}
}

// MakeVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
	return []*VirtualDNSRecord{}
}
