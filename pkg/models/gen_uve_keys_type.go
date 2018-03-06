package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeUveKeysType makes UveKeysType
// nolint
func MakeUveKeysType() *UveKeysType {
	return &UveKeysType{
		//TODO(nati): Apply default
		UveKey: []string{},
	}
}

// MakeUveKeysType makes UveKeysType
// nolint
func InterfaceToUveKeysType(i interface{}) *UveKeysType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &UveKeysType{
		//TODO(nati): Apply default
		UveKey: common.InterfaceToStringList(m["uve_key"]),
	}
}

// MakeUveKeysTypeSlice() makes a slice of UveKeysType
// nolint
func MakeUveKeysTypeSlice() []*UveKeysType {
	return []*UveKeysType{}
}

// InterfaceToUveKeysTypeSlice() makes a slice of UveKeysType
// nolint
func InterfaceToUveKeysTypeSlice(i interface{}) []*UveKeysType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*UveKeysType{}
	for _, item := range list {
		result = append(result, InterfaceToUveKeysType(item))
	}
	return result
}
