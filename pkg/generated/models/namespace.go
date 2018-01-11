package models

// Namespace

import "encoding/json"

// Namespace
type Namespace struct {
	NamespaceCidr *SubnetType    `json:"namespace_cidr"`
	ParentType    string         `json:"parent_type"`
	IDPerms       *IdPermsType   `json:"id_perms"`
	Annotations   *KeyValuePairs `json:"annotations"`
	FQName        []string       `json:"fq_name"`
	DisplayName   string         `json:"display_name"`
	Perms2        *PermType2     `json:"perms2"`
	UUID          string         `json:"uuid"`
	ParentUUID    string         `json:"parent_uuid"`
}

// String returns json representation of the object
func (model *Namespace) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNamespace makes Namespace
func MakeNamespace() *Namespace {
	return &Namespace{
		//TODO(nati): Apply default
		FQName:        []string{},
		DisplayName:   "",
		Perms2:        MakePermType2(),
		UUID:          "",
		ParentUUID:    "",
		NamespaceCidr: MakeSubnetType(),
		ParentType:    "",
		IDPerms:       MakeIdPermsType(),
		Annotations:   MakeKeyValuePairs(),
	}
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
	return []*Namespace{}
}
