package models

// JunosServicePorts

import "encoding/json"

// JunosServicePorts
type JunosServicePorts struct {
	ServicePort []string `json:"service_port,omitempty"`
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

// MakeJunosServicePortsSlice() makes a slice of JunosServicePorts
func MakeJunosServicePortsSlice() []*JunosServicePorts {
	return []*JunosServicePorts{}
}
