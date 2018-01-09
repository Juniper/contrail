package models

// User

import "encoding/json"

// User
type User struct {
	Annotations *KeyValuePairs `json:"annotations"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	Password    string         `json:"password"`
	DisplayName string         `json:"display_name"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	Perms2      *PermType2     `json:"perms2"`
	ParentType  string         `json:"parent_type"`
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
		Perms2:      MakePermType2(),
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		Password:    "",
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentUUID:  "",
	}
}

// InterfaceToUser makes User from interface
func InterfaceToUser(iData interface{}) *User {
	data := iData.(map[string]interface{})
	return &User{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Password: data["password"].(string),

		//{"description":"Domain level quota, not currently implemented","type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}

	}
}

// InterfaceToUserSlice makes a slice of User from interface
func InterfaceToUserSlice(data interface{}) []*User {
	list := data.([]interface{})
	result := MakeUserSlice()
	for _, item := range list {
		result = append(result, InterfaceToUser(item))
	}
	return result
}

// MakeUserSlice() makes a slice of User
func MakeUserSlice() []*User {
	return []*User{}
}
