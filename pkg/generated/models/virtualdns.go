package models

// VirtualDNS

import "encoding/json"

// VirtualDNS
type VirtualDNS struct {
	UUID           string          `json:"uuid"`
	ParentType     string          `json:"parent_type"`
	VirtualDNSData *VirtualDnsType `json:"virtual_DNS_data"`
	DisplayName    string          `json:"display_name"`
	Perms2         *PermType2      `json:"perms2"`
	IDPerms        *IdPermsType    `json:"id_perms"`
	Annotations    *KeyValuePairs  `json:"annotations"`
	ParentUUID     string          `json:"parent_uuid"`
	FQName         []string        `json:"fq_name"`

	VirtualDNSRecords []*VirtualDNSRecord `json:"virtual_DNS_records"`
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
		VirtualDNSData: MakeVirtualDnsType(),
		DisplayName:    "",
		Perms2:         MakePermType2(),
		UUID:           "",
		ParentType:     "",
		Annotations:    MakeKeyValuePairs(),
		ParentUUID:     "",
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
	}
}

// MakeVirtualDNSSlice() makes a slice of VirtualDNS
func MakeVirtualDNSSlice() []*VirtualDNS {
	return []*VirtualDNS{}
}
