package models

// VirtualDNSRecord

import "encoding/json"

// VirtualDNSRecord
type VirtualDNSRecord struct {
	Perms2               *PermType2            `json:"perms2"`
	ParentUUID           string                `json:"parent_uuid"`
	ParentType           string                `json:"parent_type"`
	FQName               []string              `json:"fq_name"`
	IDPerms              *IdPermsType          `json:"id_perms"`
	VirtualDNSRecordData *VirtualDnsRecordType `json:"virtual_DNS_record_data"`
	Annotations          *KeyValuePairs        `json:"annotations"`
	UUID                 string                `json:"uuid"`
	DisplayName          string                `json:"display_name"`
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
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		DisplayName:          "",
		Perms2:               MakePermType2(),
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeVirtualDNSRecordSlice() makes a slice of VirtualDNSRecord
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
	return []*VirtualDNSRecord{}
}
