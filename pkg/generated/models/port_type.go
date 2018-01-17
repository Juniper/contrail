package models

// PortType

import "encoding/json"

// PortType
type PortType struct {
	StartPort L4PortType `json:"start_port,omitempty"`
	EndPort   L4PortType `json:"end_port,omitempty"`
}

// String returns json representation of the object
func (model *PortType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePortType makes PortType
func MakePortType() *PortType {
	return &PortType{
		//TODO(nati): Apply default
		EndPort:   MakeL4PortType(),
		StartPort: MakeL4PortType(),
	}
}

// MakePortTypeSlice() makes a slice of PortType
func MakePortTypeSlice() []*PortType {
	return []*PortType{}
}
