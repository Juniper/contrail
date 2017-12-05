package models

// LinklocalServiceEntryType

import "encoding/json"

type LinklocalServiceEntryType struct {
	IPFabricServiceIP      []string `json:"ip_fabric_service_ip"`
	LinklocalServiceName   string   `json:"linklocal_service_name"`
	LinklocalServiceIP     string   `json:"linklocal_service_ip"`
	IPFabricServicePort    int      `json:"ip_fabric_service_port"`
	IPFabricDNSServiceName string   `json:"ip_fabric_DNS_service_name"`
	LinklocalServicePort   int      `json:"linklocal_service_port"`
}

func (model *LinklocalServiceEntryType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeLinklocalServiceEntryType() *LinklocalServiceEntryType {
	return &LinklocalServiceEntryType{
		//TODO(nati): Apply default
		IPFabricServiceIP:      []string{},
		LinklocalServiceName:   "",
		LinklocalServiceIP:     "",
		IPFabricServicePort:    0,
		IPFabricDNSServiceName: "",
		LinklocalServicePort:   0,
	}
}

func InterfaceToLinklocalServiceEntryType(iData interface{}) *LinklocalServiceEntryType {
	data := iData.(map[string]interface{})
	return &LinklocalServiceEntryType{
		IPFabricServiceIP: data["ip_fabric_service_ip"].([]string),

		//{"Title":"","Description":"Destination ip address to which link local traffic will forwarded","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricServiceIP","GoType":"string"},"GoName":"IPFabricServiceIP","GoType":"[]string"}
		LinklocalServiceName: data["linklocal_service_name"].(string),

		//{"Title":"","Description":"Name of the link local service. VM name resolution of this name will result in link local ip address","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceName","GoType":"string"}
		LinklocalServiceIP: data["linklocal_service_ip"].(string),

		//{"Title":"","Description":"ip address of the link local service.","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServiceIP","GoType":"string"}
		IPFabricServicePort: data["ip_fabric_service_port"].(int),

		//{"Title":"","Description":"Destination TCP port number to which link local traffic will forwarded","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricServicePort","GoType":"int"}
		IPFabricDNSServiceName: data["ip_fabric_DNS_service_name"].(string),

		//{"Title":"","Description":"DNS name to which link local service will be proxied","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPFabricDNSServiceName","GoType":"string"}
		LinklocalServicePort: data["linklocal_service_port"].(int),

		//{"Title":"","Description":"Destination TCP port number of link local service","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LinklocalServicePort","GoType":"int"}

	}
}

func InterfaceToLinklocalServiceEntryTypeSlice(data interface{}) []*LinklocalServiceEntryType {
	list := data.([]interface{})
	result := MakeLinklocalServiceEntryTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLinklocalServiceEntryType(item))
	}
	return result
}

func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
	return []*LinklocalServiceEntryType{}
}
