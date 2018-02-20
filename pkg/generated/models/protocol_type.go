package models

// ProtocolType

// ProtocolType
//proteus:generate
type ProtocolType struct {
	Protocol string `json:"protocol,omitempty"`
	Port     int    `json:"port,omitempty"`
}

// MakeProtocolType makes ProtocolType
func MakeProtocolType() *ProtocolType {
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: "",
		Port:     0,
	}
}

// MakeProtocolTypeSlice() makes a slice of ProtocolType
func MakeProtocolTypeSlice() []*ProtocolType {
	return []*ProtocolType{}
}
