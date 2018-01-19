package models

// AddressGroup

import "encoding/json"

// AddressGroup
type AddressGroup struct {
	Perms2             *PermType2      `json:"perms2,omitempty"`
	ParentUUID         string          `json:"parent_uuid,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	FQName             []string        `json:"fq_name,omitempty"`
	IDPerms            *IdPermsType    `json:"id_perms,omitempty"`
	Annotations        *KeyValuePairs  `json:"annotations,omitempty"`
	UUID               string          `json:"uuid,omitempty"`
	ParentType         string          `json:"parent_type,omitempty"`
	AddressGroupPrefix *SubnetListType `json:"address_group_prefix,omitempty"`
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
		Annotations:        MakeKeyValuePairs(),
		UUID:               "",
		ParentType:         "",
		AddressGroupPrefix: MakeSubnetListType(),
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Perms2:             MakePermType2(),
		ParentUUID:         "",
	}
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}
