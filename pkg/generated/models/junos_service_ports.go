package models

// JunosServicePorts

// JunosServicePorts
//proteus:generate
type JunosServicePorts struct {
	ServicePort []string `json:"service_port,omitempty"`
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
