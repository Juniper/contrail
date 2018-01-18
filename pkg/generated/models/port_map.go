package models

// PortMap

import "encoding/json"

// PortMap
type PortMap struct {
	DSTPort  int    `json:"dst_port,omitempty"`
	SRCPort  int    `json:"src_port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

// String returns json representation of the object
func (model *PortMap) String() string {
	b, _ := json.Marshal(model)
	return string(b)
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
