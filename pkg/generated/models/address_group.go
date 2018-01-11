package models

// AddressGroup

import "encoding/json"

// AddressGroup
type AddressGroup struct {
	IDPerms            *IdPermsType    `json:"id_perms"`
	DisplayName        string          `json:"display_name"`
	Annotations        *KeyValuePairs  `json:"annotations"`
	Perms2             *PermType2      `json:"perms2"`
	ParentUUID         string          `json:"parent_uuid"`
	AddressGroupPrefix *SubnetListType `json:"address_group_prefix"`
	UUID               string          `json:"uuid"`
	ParentType         string          `json:"parent_type"`
	FQName             []string        `json:"fq_name"`
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
		UUID:               "",
		ParentType:         "",
		FQName:             []string{},
		AddressGroupPrefix: MakeSubnetListType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		ParentUUID:         "",
		IDPerms:            MakeIdPermsType(),
	}
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}
