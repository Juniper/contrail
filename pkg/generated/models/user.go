package models

// User

import "encoding/json"

// User
type User struct {
	IDPerms     *IdPermsType   `json:"id_perms"`
	ParentType  string         `json:"parent_type"`
	Password    string         `json:"password"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	FQName      []string       `json:"fq_name"`
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
		FQName:      []string{},
		Password:    "",
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		IDPerms:     MakeIdPermsType(),
		ParentType:  "",
	}
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
	return []*User{}
}
