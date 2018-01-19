package models

// User

import "encoding/json"

// User
type User struct {
	FQName      []string       `json:"fq_name,omitempty"`
	Password    string         `json:"password,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
}

// String returns json representation of the object
func (model *User) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUser makes User
func MakeUser() *User {
	return &User{
		//TODO(nati): Apply default
		Password:    "",
		DisplayName: "",
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		IDPerms:     MakeIdPermsType(),
	}
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
	return []*User{}
}
