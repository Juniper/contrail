package models

// PortType

import "encoding/json"

// PortType
type PortType struct {
	StartPort L4PortType `json:"start_port"`
	EndPort   L4PortType `json:"end_port"`
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

// InterfaceToPortType makes PortType from interface
func InterfaceToPortType(iData interface{}) *PortType {
	data := iData.(map[string]interface{})
	return &PortType{
		EndPort: InterfaceToL4PortType(data["end_port"]),

		//{"type":"integer","minimum":-1,"maximum":65535}
		StartPort: InterfaceToL4PortType(data["start_port"]),

		//{"type":"integer","minimum":-1,"maximum":65535}

	}
}

// InterfaceToPortTypeSlice makes a slice of PortType from interface
func InterfaceToPortTypeSlice(data interface{}) []*PortType {
	list := data.([]interface{})
	result := MakePortTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToPortType(item))
	}
	return result
}

// MakePortTypeSlice() makes a slice of PortType
func MakePortTypeSlice() []*PortType {
	return []*PortType{}
}
