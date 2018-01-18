package models

// Namespace

import "encoding/json"

// Namespace
type Namespace struct {
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	NamespaceCidr *SubnetType    `json:"namespace_cidr,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
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
		NamespaceCidr: MakeSubnetType(),
		ParentUUID:    "",
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		UUID:          "",
		ParentType:    "",
		FQName:        []string{},
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
	}
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
	return []*Namespace{}
}
