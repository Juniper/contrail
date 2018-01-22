package models

// LoadbalancerListenerType

import "encoding/json"

// LoadbalancerListenerType
type LoadbalancerListenerType struct {
	DefaultTLSContainer string                   `json:"default_tls_container,omitempty"`
	Protocol            LoadbalancerProtocolType `json:"protocol,omitempty"`
	ConnectionLimit     int                      `json:"connection_limit,omitempty"`
	AdminState          bool                     `json:"admin_state"`
	SniContainers       []string                 `json:"sni_containers,omitempty"`
	ProtocolPort        int                      `json:"protocol_port,omitempty"`
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
		SniContainers:       []string{},
		ProtocolPort:        0,
		DefaultTLSContainer: "",
		Protocol:            MakeLoadbalancerProtocolType(),
		ConnectionLimit:     0,
		AdminState:          false,
	}
}

// MakeLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
	return []*LoadbalancerListenerType{}
}
