package models

// VirtualDNS

import "encoding/json"

// VirtualDNS
type VirtualDNS struct {
	Perms2         *PermType2      `json:"perms2,omitempty"`
	UUID           string          `json:"uuid,omitempty"`
	IDPerms        *IdPermsType    `json:"id_perms,omitempty"`
	FQName         []string        `json:"fq_name,omitempty"`
	DisplayName    string          `json:"display_name,omitempty"`
	Annotations    *KeyValuePairs  `json:"annotations,omitempty"`
	ParentUUID     string          `json:"parent_uuid,omitempty"`
	ParentType     string          `json:"parent_type,omitempty"`
	VirtualDNSData *VirtualDnsType `json:"virtual_DNS_data,omitempty"`

	VirtualDNSRecords []*VirtualDNSRecord `json:"virtual_DNS_records,omitempty"`
}

// String returns json representation of the object
func (model *VirtualDNS) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDNS makes VirtualDNS
func MakeVirtualDNS() *VirtualDNS {
	return &VirtualDNS{
		//TODO(nati): Apply default
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		ParentUUID:     "",
		ParentType:     "",
		VirtualDNSData: MakeVirtualDnsType(),
		FQName:         []string{},
		UUID:           "",
		IDPerms:        MakeIdPermsType(),
		Perms2:         MakePermType2(),
	}
}

// MakeVirtualDNSSlice() makes a slice of VirtualDNS
func MakeVirtualDNSSlice() []*VirtualDNS {
	return []*VirtualDNS{}
}
