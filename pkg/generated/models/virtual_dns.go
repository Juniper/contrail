package models

// VirtualDNS

import "encoding/json"

// VirtualDNS
type VirtualDNS struct {
	ParentUUID     string          `json:"parent_uuid,omitempty"`
	Annotations    *KeyValuePairs  `json:"annotations,omitempty"`
	UUID           string          `json:"uuid,omitempty"`
	DisplayName    string          `json:"display_name,omitempty"`
	Perms2         *PermType2      `json:"perms2,omitempty"`
	VirtualDNSData *VirtualDnsType `json:"virtual_DNS_data,omitempty"`
	ParentType     string          `json:"parent_type,omitempty"`
	FQName         []string        `json:"fq_name,omitempty"`
	IDPerms        *IdPermsType    `json:"id_perms,omitempty"`

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
		Annotations:    MakeKeyValuePairs(),
		UUID:           "",
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Perms2:         MakePermType2(),
		VirtualDNSData: MakeVirtualDnsType(),
	}
}

// MakeVirtualDNSSlice() makes a slice of VirtualDNS
func MakeVirtualDNSSlice() []*VirtualDNS {
	return []*VirtualDNS{}
}
