package models

// LoadbalancerListenerType

import "encoding/json"

// LoadbalancerListenerType
type LoadbalancerListenerType struct {
	ProtocolPort        int                      `json:"protocol_port"`
	DefaultTLSContainer string                   `json:"default_tls_container"`
	Protocol            LoadbalancerProtocolType `json:"protocol"`
	ConnectionLimit     int                      `json:"connection_limit"`
	AdminState          bool                     `json:"admin_state"`
	SniContainers       []string                 `json:"sni_containers"`
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
		AdminState:          false,
		SniContainers:       []string{},
		ProtocolPort:        0,
		DefaultTLSContainer: "",
		Protocol:            MakeLoadbalancerProtocolType(),
		ConnectionLimit:     0,
	}
}

// InterfaceToLoadbalancerListenerType makes LoadbalancerListenerType from interface
func InterfaceToLoadbalancerListenerType(iData interface{}) *LoadbalancerListenerType {
	data := iData.(map[string]interface{})
	return &LoadbalancerListenerType{
		SniContainers: data["sni_containers"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ProtocolPort: data["protocol_port"].(int),

		//{"type":"integer"}
		DefaultTLSContainer: data["default_tls_container"].(string),

		//{"type":"string"}
		Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),

		//{"type":"string","enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"]}
		ConnectionLimit: data["connection_limit"].(int),

		//{"type":"integer"}
		AdminState: data["admin_state"].(bool),

		//{"type":"boolean"}

	}
}

// InterfaceToLoadbalancerListenerTypeSlice makes a slice of LoadbalancerListenerType from interface
func InterfaceToLoadbalancerListenerTypeSlice(data interface{}) []*LoadbalancerListenerType {
	list := data.([]interface{})
	result := MakeLoadbalancerListenerTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerListenerType(item))
	}
	return result
}

// MakeLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
	return []*LoadbalancerListenerType{}
}
