package models

// FirewallServiceType

// FirewallServiceType
//proteus:generate
type FirewallServiceType struct {
	Protocol   string    `json:"protocol,omitempty"`
	DSTPorts   *PortType `json:"dst_ports,omitempty"`
	SRCPorts   *PortType `json:"src_ports,omitempty"`
	ProtocolID int       `json:"protocol_id,omitempty"`
}

// MakeFirewallServiceType makes FirewallServiceType
func MakeFirewallServiceType() *FirewallServiceType {
	return &FirewallServiceType{
		//TODO(nati): Apply default
		Protocol:   "",
		DSTPorts:   MakePortType(),
		SRCPorts:   MakePortType(),
		ProtocolID: 0,
	}
}

// MakeFirewallServiceTypeSlice() makes a slice of FirewallServiceType
func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
	return []*FirewallServiceType{}
}
