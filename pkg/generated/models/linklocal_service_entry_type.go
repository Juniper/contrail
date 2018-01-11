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

// MakeLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
	return []*LinklocalServiceEntryType{}
}
