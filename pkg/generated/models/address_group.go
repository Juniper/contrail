package models

// AddressGroup

import "encoding/json"

// AddressGroup
type AddressGroup struct {
	AddressGroupPrefix *SubnetListType `json:"address_group_prefix,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	UUID               string          `json:"uuid,omitempty"`
	ParentUUID         string          `json:"parent_uuid,omitempty"`
	Annotations        *KeyValuePairs  `json:"annotations,omitempty"`
	Perms2             *PermType2      `json:"perms2,omitempty"`
	ParentType         string          `json:"parent_type,omitempty"`
	FQName             []string        `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType    `json:"id_perms,omitempty"`
}

// String returns json representation of the object
func (model *AddressGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAddressGroup makes AddressGroup
func MakeAddressGroup() *AddressGroup {
	return &AddressGroup{
		//TODO(nati): Apply default
		AddressGroupPrefix: MakeSubnetListType(),
		DisplayName:        "",
		UUID:               "",
		ParentUUID:         "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
	}
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}
