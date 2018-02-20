package models

// BGPaaServiceParametersType

// BGPaaServiceParametersType
//proteus:generate
type BGPaaServiceParametersType struct {
	PortStart L4PortType `json:"port_start,omitempty"`
	PortEnd   L4PortType `json:"port_end,omitempty"`
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
