package models

// Keypair

import "encoding/json"

// Keypair
//proteus:generate
type Keypair struct {
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	Name        string         `json:"name,omitempty"`
	PublicKey   string         `json:"public_key,omitempty"`
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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		Name:        "",
		PublicKey:   "",
	}
}

// MakeKeypairSlice() makes a slice of Keypair
func MakeKeypairSlice() []*Keypair {
	return []*Keypair{}
}
