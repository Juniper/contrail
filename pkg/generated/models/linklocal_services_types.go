package models

// LinklocalServicesTypes

import "encoding/json"

type LinklocalServicesTypes struct {
	LinklocalServiceEntry []*LinklocalServiceEntryType `json:"linklocal_service_entry"`
}

func (model *LinklocalServicesTypes) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeLinklocalServicesTypes() *LinklocalServicesTypes {
	return &LinklocalServicesTypes{
		//TODO(nati): Apply default

		LinklocalServiceEntry: MakeLinklocalServiceEntryTypeSlice(),
	}
}

func InterfaceToLinklocalServicesTypes(iData interface{}) *LinklocalServicesTypes {
	data := iData.(map[string]interface{})
	return &LinklocalServicesTypes{

		LinklocalServiceEntry: InterfaceToLinklocalServiceEntryTypeSlice(data["linklocal_service_entry"]),

		//{"Title":"","Description":"List of link local services","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_fabric_DNS_service_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricDNSServiceName","GoType":"string"},"ip_fabric_service_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricServiceIP","GoType":"string"},"GoName":"IPFabricServiceIP","GoType":"[]string"},"ip_fabric_service_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricServicePort","GoType":"int"},"linklocal_service_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceIP","GoType":"string"},"linklocal_service_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceName","GoType":"string"},"linklocal_service_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServicePort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/LinklocalServiceEntryType","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceEntry","GoType":"LinklocalServiceEntryType"},"GoName":"LinklocalServiceEntry","GoType":"[]*LinklocalServiceEntryType"}

	}
}

func InterfaceToLinklocalServicesTypesSlice(data interface{}) []*LinklocalServicesTypes {
	list := data.([]interface{})
	result := MakeLinklocalServicesTypesSlice()
	for _, item := range list {
		result = append(result, InterfaceToLinklocalServicesTypes(item))
	}
	return result
}

func MakeLinklocalServicesTypesSlice() []*LinklocalServicesTypes {
	return []*LinklocalServicesTypes{}
}
