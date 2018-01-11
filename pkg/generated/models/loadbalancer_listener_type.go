package models

// LoadbalancerListenerType

import "encoding/json"

// LoadbalancerListenerType
type LoadbalancerListenerType struct {
	DefaultTLSContainer string                   `json:"default_tls_container"`
	Protocol            LoadbalancerProtocolType `json:"protocol"`
	ConnectionLimit     int                      `json:"connection_limit"`
	AdminState          bool                     `json:"admin_state"`
	SniContainers       []string                 `json:"sni_containers"`
	ProtocolPort        int                      `json:"protocol_port"`
}

// String returns json representation of the object
func (model *LoadbalancerListenerType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerListenerType makes LoadbalancerListenerType
func MakeLoadbalancerListenerType() *LoadbalancerListenerType {
	return &LoadbalancerListenerType{
		//TODO(nati): Apply default
		ConnectionLimit:     0,
		AdminState:          false,
		SniContainers:       []string{},
		ProtocolPort:        0,
		DefaultTLSContainer: "",
		Protocol:            MakeLoadbalancerProtocolType(),
	}
}

// MakeLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
	return []*LoadbalancerListenerType{}
}
