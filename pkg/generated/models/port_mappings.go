package models

// PortMappings

import "encoding/json"

// PortMappings
type PortMappings struct {
	PortMappings []*PortMap `json:"port_mappings"`
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

// InterfaceToPortMappings makes PortMappings from interface
func InterfaceToPortMappings(iData interface{}) *PortMappings {
	data := iData.(map[string]interface{})
	return &PortMappings{

		PortMappings: InterfaceToPortMapSlice(data["port_mappings"]),

		//{"type":"array","item":{"type":"object","properties":{"dst_port":{"type":"integer"},"protocol":{"type":"string"},"src_port":{"type":"integer"}}}}

	}
}

// InterfaceToPortMappingsSlice makes a slice of PortMappings from interface
func InterfaceToPortMappingsSlice(data interface{}) []*PortMappings {
	list := data.([]interface{})
	result := MakePortMappingsSlice()
	for _, item := range list {
		result = append(result, InterfaceToPortMappings(item))
	}
	return result
}

// MakePortMappingsSlice() makes a slice of PortMappings
func MakePortMappingsSlice() []*PortMappings {
	return []*PortMappings{}
}
