package models

// LinklocalServicesTypes

// LinklocalServicesTypes
//proteus:generate
type LinklocalServicesTypes struct {
	LinklocalServiceEntry []*LinklocalServiceEntryType `json:"linklocal_service_entry,omitempty"`
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
