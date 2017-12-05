package models

// LoadbalancerListenerType

import "encoding/json"

type LoadbalancerListenerType struct {
	DefaultTLSContainer string                   `json:"default_tls_container"`
	Protocol            LoadbalancerProtocolType `json:"protocol"`
	ConnectionLimit     int                      `json:"connection_limit"`
	AdminState          bool                     `json:"admin_state"`
	SniContainers       []string                 `json:"sni_containers"`
	ProtocolPort        int                      `json:"protocol_port"`
}

func (model *LoadbalancerListenerType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

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

func InterfaceToLoadbalancerListenerType(iData interface{}) *LoadbalancerListenerType {
	data := iData.(map[string]interface{})
	return &LoadbalancerListenerType{
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool"}
		SniContainers: data["sni_containers"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SniContainers","GoType":"string"},"GoName":"SniContainers","GoType":"[]string"}
		ProtocolPort: data["protocol_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolPort","GoType":"int"}
		DefaultTLSContainer: data["default_tls_container"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DefaultTLSContainer","GoType":"string"}
		Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/LoadbalancerProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"LoadbalancerProtocolType"}
		ConnectionLimit: data["connection_limit"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ConnectionLimit","GoType":"int"}

	}
}

func InterfaceToLoadbalancerListenerTypeSlice(data interface{}) []*LoadbalancerListenerType {
	list := data.([]interface{})
	result := MakeLoadbalancerListenerTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerListenerType(item))
	}
	return result
}

func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
	return []*LoadbalancerListenerType{}
}
