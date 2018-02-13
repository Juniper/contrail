package models

// BGPaaServiceParametersType

import "encoding/json"

// BGPaaServiceParametersType
//proteus:generate
type BGPaaServiceParametersType struct {
	PortStart L4PortType `json:"port_start,omitempty"`
	PortEnd   L4PortType `json:"port_end,omitempty"`
}

// String returns json representation of the object
func (model *BGPaaServiceParametersType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
func MakeBGPaaServiceParametersType() *BGPaaServiceParametersType {
	return &BGPaaServiceParametersType{
		//TODO(nati): Apply default
		PortStart: MakeL4PortType(),
		PortEnd:   MakeL4PortType(),
	}
}

// MakeBGPaaServiceParametersTypeSlice() makes a slice of BGPaaServiceParametersType
func MakeBGPaaServiceParametersTypeSlice() []*BGPaaServiceParametersType {
	return []*BGPaaServiceParametersType{}
}
