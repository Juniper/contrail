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

// InterfaceToLinklocalServicesTypes makes LinklocalServicesTypes from interface
func InterfaceToLinklocalServicesTypes(iData interface{}) *LinklocalServicesTypes {
	data := iData.(map[string]interface{})
	return &LinklocalServicesTypes{

		LinklocalServiceEntry: InterfaceToLinklocalServiceEntryTypeSlice(data["linklocal_service_entry"]),

		//{"description":"List of link local services","type":"array","item":{"type":"object","properties":{"ip_fabric_DNS_service_name":{"type":"string"},"ip_fabric_service_ip":{"type":"array","item":{"type":"string"}},"ip_fabric_service_port":{"type":"integer"},"linklocal_service_ip":{"type":"string"},"linklocal_service_name":{"type":"string"},"linklocal_service_port":{"type":"integer"}}}}

	}
}

// InterfaceToLinklocalServicesTypesSlice makes a slice of LinklocalServicesTypes from interface
func InterfaceToLinklocalServicesTypesSlice(data interface{}) []*LinklocalServicesTypes {
	list := data.([]interface{})
	result := MakeLinklocalServicesTypesSlice()
	for _, item := range list {
		result = append(result, InterfaceToLinklocalServicesTypes(item))
	}
	return result
}

// MakeLinklocalServicesTypesSlice() makes a slice of LinklocalServicesTypes
func MakeLinklocalServicesTypesSlice() []*LinklocalServicesTypes {
	return []*LinklocalServicesTypes{}
}
