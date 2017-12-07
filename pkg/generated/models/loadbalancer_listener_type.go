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

//  parents relation object

// String returns json representation of the object
func (model *LoadbalancerListenerType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerListenerType makes LoadbalancerListenerType
func MakeLoadbalancerListenerType() *LoadbalancerListenerType {
	return &LoadbalancerListenerType{
		//TODO(nati): Apply default
		DefaultTLSContainer: "",
		Protocol:            MakeLoadbalancerProtocolType(),
		ConnectionLimit:     0,
		AdminState:          false,
		SniContainers:       []string{},
		ProtocolPort:        0,
	}
}

// InterfaceToLoadbalancerListenerType makes LoadbalancerListenerType from interface
func InterfaceToLoadbalancerListenerType(iData interface{}) *LoadbalancerListenerType {
	data := iData.(map[string]interface{})
	return &LoadbalancerListenerType{
		DefaultTLSContainer: data["default_tls_container"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"DefaultTLSContainer","GoType":"string","GoPremitive":true}
		Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/LoadbalancerProtocolType","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"LoadbalancerProtocolType","GoPremitive":false}
		ConnectionLimit: data["connection_limit"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ConnectionLimit","GoType":"int","GoPremitive":true}
		AdminState: data["admin_state"].(bool),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AdminState","GoType":"bool","GoPremitive":true}
		SniContainers: data["sni_containers"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SniContainers","GoType":"string","GoPremitive":true},"GoName":"SniContainers","GoType":"[]string","GoPremitive":true}
		ProtocolPort: data["protocol_port"].(int),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolPort","GoType":"int","GoPremitive":true}

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
