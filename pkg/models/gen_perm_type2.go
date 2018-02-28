package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakePermType2 makes PermType2
// nolint
func MakePermType2() *PermType2 {
	return &PermType2{
		//TODO(nati): Apply default
		Owner:        "",
		OwnerAccess:  0,
		GlobalAccess: 0,

		Share: MakeShareTypeSlice(),
	}
}

// MakePermType2 makes PermType2
// nolint
func InterfaceToPermType2(i interface{}) *PermType2 {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &PermType2{
		//TODO(nati): Apply default
		Owner:        common.InterfaceToString(m["owner"]),
		OwnerAccess:  common.InterfaceToInt64(m["owner_access"]),
		GlobalAccess: common.InterfaceToInt64(m["global_access"]),

		Share: InterfaceToShareTypeSlice(m["share"]),
	}
}

// MakePermType2Slice() makes a slice of PermType2
// nolint
func MakePermType2Slice() []*PermType2 {
	return []*PermType2{}
}

// InterfaceToPermType2Slice() makes a slice of PermType2
// nolint
func InterfaceToPermType2Slice(i interface{}) []*PermType2 {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*PermType2{}
	for _, item := range list {
		result = append(result, InterfaceToPermType2(item))
	}
	return result
}
