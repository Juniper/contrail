package models

// Keypair

import "encoding/json"

// Keypair
type Keypair struct {
	PublicKey   string         `json:"public_key,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	Name        string         `json:"name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
}

// String returns json representation of the object
func (model *Keypair) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeKeypair makes Keypair
func MakeKeypair() *Keypair {
	return &Keypair{
		//TODO(nati): Apply default
		Name:        "",
		IDPerms:     MakeIdPermsType(),
		ParentUUID:  "",
		ParentType:  "",
		UUID:        "",
		PublicKey:   "",
		FQName:      []string{},
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeKeypairSlice() makes a slice of Keypair
func MakeKeypairSlice() []*Keypair {
	return []*Keypair{}
}
