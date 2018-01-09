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
		Protocol: "",
		DSTPort:  0,
		SRCPort:  0,
	}
}

// InterfaceToPortMap makes PortMap from interface
func InterfaceToPortMap(iData interface{}) *PortMap {
	data := iData.(map[string]interface{})
	return &PortMap{
		SRCPort: data["src_port"].(int),

		//{"type":"integer"}
		Protocol: data["protocol"].(string),

		//{"type":"string"}
		DSTPort: data["dst_port"].(int),

		//{"type":"integer"}

	}
}

// InterfaceToPortMapSlice makes a slice of PortMap from interface
func InterfaceToPortMapSlice(data interface{}) []*PortMap {
	list := data.([]interface{})
	result := MakePortMapSlice()
	for _, item := range list {
		result = append(result, InterfaceToPortMap(item))
	}
	return result
}

// MakePortMapSlice() makes a slice of PortMap
func MakePortMapSlice() []*PortMap {
	return []*PortMap{}
}
