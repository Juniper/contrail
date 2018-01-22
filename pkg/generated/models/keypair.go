package models

// Keypair

import "encoding/json"

// Keypair
type Keypair struct {
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	PublicKey   string         `json:"public_key,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	Name        string         `json:"name,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
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
		DisplayName: "",
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		PublicKey:   "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeKeypairSlice() makes a slice of Keypair
func MakeKeypairSlice() []*Keypair {
	return []*Keypair{}
}
