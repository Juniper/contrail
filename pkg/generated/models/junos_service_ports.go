package models

// JunosServicePorts

import "encoding/json"

// JunosServicePorts
type JunosServicePorts struct {
	ServicePort []string `json:"service_port"`
}

// String returns json representation of the object
func (model *JunosServicePorts) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeJunosServicePorts makes JunosServicePorts
func MakeJunosServicePorts() *JunosServicePorts {
	return &JunosServicePorts{
		//TODO(nati): Apply default
		ServicePort: []string{},
	}
}

// InterfaceToJunosServicePorts makes JunosServicePorts from interface
func InterfaceToJunosServicePorts(iData interface{}) *JunosServicePorts {
	data := iData.(map[string]interface{})
	return &JunosServicePorts{
		ServicePort: data["service_port"].([]string),

		//{"type":"array","item":{"type":"string"}}

	}
}

// InterfaceToJunosServicePortsSlice makes a slice of JunosServicePorts from interface
func InterfaceToJunosServicePortsSlice(data interface{}) []*JunosServicePorts {
	list := data.([]interface{})
	result := MakeJunosServicePortsSlice()
	for _, item := range list {
		result = append(result, InterfaceToJunosServicePorts(item))
	}
	return result
}

// MakeJunosServicePortsSlice() makes a slice of JunosServicePorts
func MakeJunosServicePortsSlice() []*JunosServicePorts {
	return []*JunosServicePorts{}
}
