package models

// User

import "encoding/json"

// User
type User struct {
	Perms2      *PermType2     `json:"perms2,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	Password    string         `json:"password,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
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
		IDPerms:     MakeIdPermsType(),
		Password:    "",
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		Perms2:      MakePermType2(),
		FQName:      []string{},
	}
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
	return []*User{}
}
