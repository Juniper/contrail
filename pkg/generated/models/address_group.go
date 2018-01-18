package models

// AddressGroup

import "encoding/json"

// AddressGroup
type AddressGroup struct {
	FQName             []string        `json:"fq_name,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs  `json:"annotations,omitempty"`
	AddressGroupPrefix *SubnetListType `json:"address_group_prefix,omitempty"`
	UUID               string          `json:"uuid,omitempty"`
	ParentType         string          `json:"parent_type,omitempty"`
	ParentUUID         string          `json:"parent_uuid,omitempty"`
	IDPerms            *IdPermsType    `json:"id_perms,omitempty"`
	Perms2             *PermType2      `json:"perms2,omitempty"`
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
		UUID:               "",
		ParentType:         "",
		FQName:             []string{},
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		ParentUUID:         "",
		IDPerms:            MakeIdPermsType(),
		Perms2:             MakePermType2(),
	}
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}
