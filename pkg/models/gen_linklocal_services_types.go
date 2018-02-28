package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLinklocalServicesTypes makes LinklocalServicesTypes
// nolint
func MakeLinklocalServicesTypes() *LinklocalServicesTypes {
	return &LinklocalServicesTypes{
		//TODO(nati): Apply default

		LinklocalServiceEntry: MakeLinklocalServiceEntryTypeSlice(),
	}
}

// MakeLinklocalServicesTypes makes LinklocalServicesTypes
// nolint
func InterfaceToLinklocalServicesTypes(i interface{}) *LinklocalServicesTypes {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LinklocalServicesTypes{
		//TODO(nati): Apply default

		LinklocalServiceEntry: InterfaceToLinklocalServiceEntryTypeSlice(m["linklocal_service_entry"]),
	}
}

// MakeLinklocalServicesTypesSlice() makes a slice of LinklocalServicesTypes
// nolint
func MakeLinklocalServicesTypesSlice() []*LinklocalServicesTypes {
	return []*LinklocalServicesTypes{}
}

// InterfaceToLinklocalServicesTypesSlice() makes a slice of LinklocalServicesTypes
// nolint
func InterfaceToLinklocalServicesTypesSlice(i interface{}) []*LinklocalServicesTypes {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LinklocalServicesTypes{}
	for _, item := range list {
		result = append(result, InterfaceToLinklocalServicesTypes(item))
	}
	return result
}
