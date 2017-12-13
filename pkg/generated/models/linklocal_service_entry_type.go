package models

// LinklocalServiceEntryType

import "encoding/json"

// LinklocalServiceEntryType
type LinklocalServiceEntryType struct {
	LinklocalServicePort   int      `json:"linklocal_service_port"`
	IPFabricServiceIP      []string `json:"ip_fabric_service_ip"`
	LinklocalServiceName   string   `json:"linklocal_service_name"`
	LinklocalServiceIP     string   `json:"linklocal_service_ip"`
	IPFabricServicePort    int      `json:"ip_fabric_service_port"`
	IPFabricDNSServiceName string   `json:"ip_fabric_DNS_service_name"`
}

// String returns json representation of the object
func (model *LinklocalServiceEntryType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLinklocalServiceEntryType makes LinklocalServiceEntryType
func MakeLinklocalServiceEntryType() *LinklocalServiceEntryType {
	return &LinklocalServiceEntryType{
		//TODO(nati): Apply default
		IPFabricDNSServiceName: "",
		LinklocalServicePort:   0,
		IPFabricServiceIP:      []string{},
		LinklocalServiceName:   "",
		LinklocalServiceIP:     "",
		IPFabricServicePort:    0,
	}
}

// InterfaceToLinklocalServiceEntryType makes LinklocalServiceEntryType from interface
func InterfaceToLinklocalServiceEntryType(iData interface{}) *LinklocalServiceEntryType {
	data := iData.(map[string]interface{})
	return &LinklocalServiceEntryType{
		LinklocalServiceName: data["linklocal_service_name"].(string),

		//{"description":"Name of the link local service. VM name resolution of this name will result in link local ip address","type":"string"}
		LinklocalServiceIP: data["linklocal_service_ip"].(string),

		//{"description":"ip address of the link local service.","type":"string"}
		IPFabricServicePort: data["ip_fabric_service_port"].(int),

		//{"description":"Destination TCP port number to which link local traffic will forwarded","type":"integer"}
		IPFabricDNSServiceName: data["ip_fabric_DNS_service_name"].(string),

		//{"description":"DNS name to which link local service will be proxied","type":"string"}
		LinklocalServicePort: data["linklocal_service_port"].(int),

		//{"description":"Destination TCP port number of link local service","type":"integer"}
		IPFabricServiceIP: data["ip_fabric_service_ip"].([]string),

		//{"description":"Destination ip address to which link local traffic will forwarded","type":"array","item":{"type":"string"}}

	}
}

// InterfaceToLinklocalServiceEntryTypeSlice makes a slice of LinklocalServiceEntryType from interface
func InterfaceToLinklocalServiceEntryTypeSlice(data interface{}) []*LinklocalServiceEntryType {
	list := data.([]interface{})
	result := MakeLinklocalServiceEntryTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLinklocalServiceEntryType(item))
	}
	return result
}

// MakeLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
	return []*LinklocalServiceEntryType{}
}
