package models

// AddressGroup

import "encoding/json"

// AddressGroup
type AddressGroup struct {
	UUID               string          `json:"uuid,omitempty"`
	ParentType         string          `json:"parent_type,omitempty"`
	AddressGroupPrefix *SubnetListType `json:"address_group_prefix,omitempty"`
	Perms2             *PermType2      `json:"perms2,omitempty"`
	IDPerms            *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs  `json:"annotations,omitempty"`
	ParentUUID         string          `json:"parent_uuid,omitempty"`
	FQName             []string        `json:"fq_name,omitempty"`
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
		Perms2:             MakePermType2(),
		UUID:               "",
		ParentType:         "",
		ParentUUID:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
	}
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}
