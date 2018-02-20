package models

// PortMap

// PortMap
//proteus:generate
type PortMap struct {
	SRCPort  int    `json:"src_port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	DSTPort  int    `json:"dst_port,omitempty"`
}

// MakePortMap makes PortMap
func MakePortMap() *PortMap {
	return &PortMap{
		//TODO(nati): Apply default
		SRCPort:  0,
		Protocol: "",
		DSTPort:  0,
	}
}

// MakePortMapSlice() makes a slice of PortMap
func MakePortMapSlice() []*PortMap {
	return []*PortMap{}
}
