package models

// ProtocolType

import "encoding/json"

// ProtocolType
type ProtocolType struct {
	Protocol string `json:"protocol,omitempty"`
	Port     int    `json:"port,omitempty"`
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
		Port:     0,
		Protocol: "",
	}
}

// MakeProtocolTypeSlice() makes a slice of ProtocolType
func MakeProtocolTypeSlice() []*ProtocolType {
	return []*ProtocolType{}
}
