package models

// PortMap

import "encoding/json"

// PortMap
type PortMap struct {
	SRCPort  int    `json:"src_port"`
	Protocol string `json:"protocol"`
	DSTPort  int    `json:"dst_port"`
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
