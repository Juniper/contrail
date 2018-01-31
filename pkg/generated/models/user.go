package models

// User

import "encoding/json"

// User
type User struct {
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Password    string         `json:"password,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
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
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		Password:    "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
	}
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
	return []*User{}
}
