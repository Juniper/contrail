package models

// Namespace

import "encoding/json"

// Namespace
type Namespace struct {
	UUID          string         `json:"uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	NamespaceCidr *SubnetType    `json:"namespace_cidr,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
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
		UUID:          "",
		ParentType:    "",
		IDPerms:       MakeIdPermsType(),
		FQName:        []string{},
		DisplayName:   "",
		NamespaceCidr: MakeSubnetType(),
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		ParentUUID:    "",
	}
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
	return []*Namespace{}
}
