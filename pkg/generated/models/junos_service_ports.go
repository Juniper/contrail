package models


// MakeJunosServicePorts makes JunosServicePorts
func MakeJunosServicePorts() *JunosServicePorts{
    return &JunosServicePorts{
    //TODO(nati): Apply default
    ServicePort: []string{},
        
    }
}

// MakeJunosServicePortsSlice() makes a slice of JunosServicePorts
func MakeJunosServicePortsSlice() []*JunosServicePorts {
    return []*JunosServicePorts{}
}


