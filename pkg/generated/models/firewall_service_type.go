package models

// FirewallServiceType

import "encoding/json"

// FirewallServiceType
type FirewallServiceType struct {
	SRCPorts   *PortType `json:"src_ports"`
	ProtocolID int       `json:"protocol_id"`
	Protocol   string    `json:"protocol"`
	DSTPorts   *PortType `json:"dst_ports"`
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
		SRCPorts:   MakePortType(),
		ProtocolID: 0,
		Protocol:   "",
		DSTPorts:   MakePortType(),
	}
}

// MakeFirewallServiceTypeSlice() makes a slice of FirewallServiceType
func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
	return []*FirewallServiceType{}
}
