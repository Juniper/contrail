package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeUserCredentials makes UserCredentials
// nolint
func MakeUserCredentials() *UserCredentials {
	return &UserCredentials{
		//TODO(nati): Apply default
		Username: "",
		Password: "",
	}
}

// MakeUserCredentials makes UserCredentials
// nolint
func InterfaceToUserCredentials(i interface{}) *UserCredentials {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &UserCredentials{
		//TODO(nati): Apply default
		Username: common.InterfaceToString(m["username"]),
		Password: common.InterfaceToString(m["password"]),
	}
}

// MakeUserCredentialsSlice() makes a slice of UserCredentials
// nolint
func MakeUserCredentialsSlice() []*UserCredentials {
	return []*UserCredentials{}
}

// InterfaceToUserCredentialsSlice() makes a slice of UserCredentials
// nolint
func InterfaceToUserCredentialsSlice(i interface{}) []*UserCredentials {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*UserCredentials{}
	for _, item := range list {
		result = append(result, InterfaceToUserCredentials(item))
	}
	return result
}
