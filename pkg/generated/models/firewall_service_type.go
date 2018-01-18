package models

// FirewallServiceType

import "encoding/json"

// FirewallServiceType
type FirewallServiceType struct {
	DSTPorts   *PortType `json:"dst_ports,omitempty"`
	SRCPorts   *PortType `json:"src_ports,omitempty"`
	ProtocolID int       `json:"protocol_id,omitempty"`
	Protocol   string    `json:"protocol,omitempty"`
}

// String returns json representation of the object
func (model *FirewallServiceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallServiceType makes FirewallServiceType
func MakeFirewallServiceType() *FirewallServiceType {
	return &FirewallServiceType{
		//TODO(nati): Apply default
		DSTPorts:   MakePortType(),
		SRCPorts:   MakePortType(),
		ProtocolID: 0,
		Protocol:   "",
	}
}

// MakeFirewallServiceTypeSlice() makes a slice of FirewallServiceType
func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
	return []*FirewallServiceType{}
}
