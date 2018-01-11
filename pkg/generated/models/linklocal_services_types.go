package models

// LinklocalServicesTypes

import "encoding/json"

// LinklocalServicesTypes
type LinklocalServicesTypes struct {
	LinklocalServiceEntry []*LinklocalServiceEntryType `json:"linklocal_service_entry"`
}

// String returns json representation of the object
func (model *LinklocalServicesTypes) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLinklocalServicesTypes makes LinklocalServicesTypes
func MakeLinklocalServicesTypes() *LinklocalServicesTypes {
	return &LinklocalServicesTypes{
		//TODO(nati): Apply default

		LinklocalServiceEntry: MakeLinklocalServiceEntryTypeSlice(),
	}
}

// MakeLinklocalServicesTypesSlice() makes a slice of LinklocalServicesTypes
func MakeLinklocalServicesTypesSlice() []*LinklocalServicesTypes {
	return []*LinklocalServicesTypes{}
}
