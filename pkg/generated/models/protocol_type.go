package models

// ProtocolType

import "encoding/json"

// ProtocolType
type ProtocolType struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

// String returns json representation of the object
func (model *ProtocolType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeProtocolType makes ProtocolType
func MakeProtocolType() *ProtocolType {
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: "",
		Port:     0,
	}
}

// InterfaceToProtocolType makes ProtocolType from interface
func InterfaceToProtocolType(iData interface{}) *ProtocolType {
	data := iData.(map[string]interface{})
	return &ProtocolType{
		Protocol: data["protocol"].(string),

		//{"type":"string"}
		Port: data["port"].(int),

		//{"type":"integer"}

	}
}

// InterfaceToProtocolTypeSlice makes a slice of ProtocolType from interface
func InterfaceToProtocolTypeSlice(data interface{}) []*ProtocolType {
	list := data.([]interface{})
	result := MakeProtocolTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToProtocolType(item))
	}
	return result
}

// MakeProtocolTypeSlice() makes a slice of ProtocolType
func MakeProtocolTypeSlice() []*ProtocolType {
	return []*ProtocolType{}
}
