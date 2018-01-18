package models

// AddressGroup

import "encoding/json"

// AddressGroup
type AddressGroup struct {
	IDPerms            *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	Annotations        *KeyValuePairs  `json:"annotations,omitempty"`
	Perms2             *PermType2      `json:"perms2,omitempty"`
	AddressGroupPrefix *SubnetListType `json:"address_group_prefix,omitempty"`
	ParentUUID         string          `json:"parent_uuid,omitempty"`
	ParentType         string          `json:"parent_type,omitempty"`
	FQName             []string        `json:"fq_name,omitempty"`
	UUID               string          `json:"uuid,omitempty"`
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
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		AddressGroupPrefix: MakeSubnetListType(),
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		UUID:               "",
	}
}

// MakeAddressGroupSlice() makes a slice of AddressGroup
func MakeAddressGroupSlice() []*AddressGroup {
	return []*AddressGroup{}
}
