package models

// LocalLinkConnection

// LocalLinkConnection
//proteus:generate
type LocalLinkConnection struct {
	SwitchID   string `json:"switch_id,omitempty"`
	PortID     string `json:"port_id,omitempty"`
	SwitchInfo string `json:"switch_info,omitempty"`
}

// MakeLocalLinkConnection makes LocalLinkConnection
func MakeLocalLinkConnection() *LocalLinkConnection {
	return &LocalLinkConnection{
		//TODO(nati): Apply default
		SwitchID:   "",
		PortID:     "",
		SwitchInfo: "",
	}
}

// MakeLocalLinkConnectionSlice() makes a slice of LocalLinkConnection
func MakeLocalLinkConnectionSlice() []*LocalLinkConnection {
	return []*LocalLinkConnection{}
}
