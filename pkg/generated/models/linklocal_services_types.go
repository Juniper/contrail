package models

// LinklocalServicesTypes

import "encoding/json"

// LinklocalServicesTypes
type LinklocalServicesTypes struct {
	LinklocalServiceEntry []*LinklocalServiceEntryType `json:"linklocal_service_entry"`
}

//  parents relation object

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

		//{"Title":"","Description":"List of link local services","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_fabric_DNS_service_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricDNSServiceName","GoType":"string","GoPremitive":true},"ip_fabric_service_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricServiceIP","GoType":"string","GoPremitive":true},"GoName":"IPFabricServiceIP","GoType":"[]string","GoPremitive":true},"ip_fabric_service_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricServicePort","GoType":"int","GoPremitive":true},"linklocal_service_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceIP","GoType":"string","GoPremitive":true},"linklocal_service_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceName","GoType":"string","GoPremitive":true},"linklocal_service_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServicePort","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/LinklocalServiceEntryType","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceEntry","GoType":"LinklocalServiceEntryType","GoPremitive":false},"GoName":"LinklocalServiceEntry","GoType":"[]*LinklocalServiceEntryType","GoPremitive":true}

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
