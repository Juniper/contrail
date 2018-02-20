package models

// PortMappings

// PortMappings
//proteus:generate
type PortMappings struct {
	PortMappings []*PortMap `json:"port_mappings,omitempty"`
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
