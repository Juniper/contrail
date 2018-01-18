package models

// LinklocalServiceEntryType

import "encoding/json"

// LinklocalServiceEntryType
type LinklocalServiceEntryType struct {
	IPFabricDNSServiceName string   `json:"ip_fabric_DNS_service_name,omitempty"`
	LinklocalServicePort   int      `json:"linklocal_service_port,omitempty"`
	IPFabricServiceIP      []string `json:"ip_fabric_service_ip,omitempty"`
	LinklocalServiceName   string   `json:"linklocal_service_name,omitempty"`
	LinklocalServiceIP     string   `json:"linklocal_service_ip,omitempty"`
	IPFabricServicePort    int      `json:"ip_fabric_service_port,omitempty"`
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
		IPFabricServiceIP:      []string{},
		LinklocalServiceName:   "",
		LinklocalServiceIP:     "",
		IPFabricServicePort:    0,
		IPFabricDNSServiceName: "",
		LinklocalServicePort:   0,
	}
}

// MakeLinklocalServiceEntryTypeSlice() makes a slice of LinklocalServiceEntryType
func MakeLinklocalServiceEntryTypeSlice() []*LinklocalServiceEntryType {
	return []*LinklocalServiceEntryType{}
}
