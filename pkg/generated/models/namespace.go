package models

// Namespace

import "encoding/json"

// Namespace
type Namespace struct {
	Perms2        *PermType2     `json:"perms2,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	NamespaceCidr *SubnetType    `json:"namespace_cidr,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
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
		ParentUUID:    "",
		ParentType:    "",
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		DisplayName:   "",
		NamespaceCidr: MakeSubnetType(),
		UUID:          "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
	}
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
	return []*Namespace{}
}
