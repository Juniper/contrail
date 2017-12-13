package models

// UserCredentials

import "encoding/json"

// UserCredentials
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// String returns json representation of the object
func (model *UserCredentials) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUserCredentials makes UserCredentials
func MakeUserCredentials() *UserCredentials {
	return &UserCredentials{
		//TODO(nati): Apply default
		Username: "",
		Password: "",
	}
}

// InterfaceToUserCredentials makes UserCredentials from interface
func InterfaceToUserCredentials(iData interface{}) *UserCredentials {
	data := iData.(map[string]interface{})
	return &UserCredentials{
		Username: data["username"].(string),

		//{"type":"string"}
		Password: data["password"].(string),

		//{"type":"string"}

	}
}

// InterfaceToUserCredentialsSlice makes a slice of UserCredentials from interface
func InterfaceToUserCredentialsSlice(data interface{}) []*UserCredentials {
	list := data.([]interface{})
	result := MakeUserCredentialsSlice()
	for _, item := range list {
		result = append(result, InterfaceToUserCredentials(item))
	}
	return result
}

// MakeUserCredentialsSlice() makes a slice of UserCredentials
func MakeUserCredentialsSlice() []*UserCredentials {
	return []*UserCredentials{}
}
