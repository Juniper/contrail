package models

// User

import "encoding/json"

// User
type User struct {
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Password    string         `json:"password,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
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
		ParentUUID:  "",
		FQName:      []string{},
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Password:    "",
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
	}
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
	return []*User{}
}
