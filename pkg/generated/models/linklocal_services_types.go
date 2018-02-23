package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeLinklocalServicesTypes makes LinklocalServicesTypes
func MakeLinklocalServicesTypes() *LinklocalServicesTypes {
	return &LinklocalServicesTypes{
		//TODO(nati): Apply default

		LinklocalServiceEntry: MakeLinklocalServiceEntryTypeSlice(),
	}
}

// MakeLinklocalServicesTypes makes LinklocalServicesTypes
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
func MakeLinklocalServicesTypesSlice() []*LinklocalServicesTypes {
	return []*LinklocalServicesTypes{}
}

// InterfaceToLinklocalServicesTypesSlice() makes a slice of LinklocalServicesTypes
func InterfaceToLinklocalServicesTypesSlice(i interface{}) []*LinklocalServicesTypes {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LinklocalServicesTypes{}
	for _, item := range list {
		result = append(result, InterfaceToLinklocalServicesTypes(item))
	}
	return result
}
