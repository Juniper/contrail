package models

// PortMappings

import "encoding/json"

// PortMappings
type PortMappings struct {
	PortMappings []*PortMap `json:"port_mappings,omitempty"`
}

// String returns json representation of the object
func (model *PortMappings) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePortMappings makes PortMappings
func MakePortMappings() *PortMappings {
	return &PortMappings{
		//TODO(nati): Apply default

		PortMappings: MakePortMapSlice(),
	}
}

// MakePortMappingsSlice() makes a slice of PortMappings
func MakePortMappingsSlice() []*PortMappings {
	return []*PortMappings{}
}
