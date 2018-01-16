package models

// UserCredentials

import "encoding/json"

// UserCredentials
type UserCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
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
		Password: "",
		Username: "",
	}
}

// MakeUserCredentialsSlice() makes a slice of UserCredentials
func MakeUserCredentialsSlice() []*UserCredentials {
	return []*UserCredentials{}
}
