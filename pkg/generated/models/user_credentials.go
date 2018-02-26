package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeUserCredentials makes UserCredentials
func MakeUserCredentials() *UserCredentials {
	return &UserCredentials{
		//TODO(nati): Apply default
		Username: "",
		Password: "",
	}
}

// MakeUserCredentials makes UserCredentials
func InterfaceToUserCredentials(i interface{}) *UserCredentials {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &UserCredentials{
		//TODO(nati): Apply default
		Username: schema.InterfaceToString(m["username"]),
		Password: schema.InterfaceToString(m["password"]),
	}
}

// MakeUserCredentialsSlice() makes a slice of UserCredentials
func MakeUserCredentialsSlice() []*UserCredentials {
	return []*UserCredentials{}
}

// InterfaceToUserCredentialsSlice() makes a slice of UserCredentials
func InterfaceToUserCredentialsSlice(i interface{}) []*UserCredentials {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*UserCredentials{}
	for _, item := range list {
		result = append(result, InterfaceToUserCredentials(item))
	}
	return result
}
