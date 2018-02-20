package models

// PortType

// PortType
//proteus:generate
type PortType struct {
	EndPort   L4PortType `json:"end_port,omitempty"`
	StartPort L4PortType `json:"start_port,omitempty"`
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
