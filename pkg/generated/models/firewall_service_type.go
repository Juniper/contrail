package models

// FirewallServiceType

import "encoding/json"

// FirewallServiceType
type FirewallServiceType struct {
	Protocol   string    `json:"protocol"`
	DSTPorts   *PortType `json:"dst_ports"`
	SRCPorts   *PortType `json:"src_ports"`
	ProtocolID int       `json:"protocol_id"`
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
		ProtocolID: 0,
		Protocol:   "",
		DSTPorts:   MakePortType(),
		SRCPorts:   MakePortType(),
	}
}

// InterfaceToFirewallServiceType makes FirewallServiceType from interface
func InterfaceToFirewallServiceType(iData interface{}) *FirewallServiceType {
	data := iData.(map[string]interface{})
	return &FirewallServiceType{
		Protocol: data["protocol"].(string),

		//{"description":"Layer 4 protocol in ip packet","type":"string"}
		DSTPorts: InterfaceToPortType(data["dst_ports"]),

		//{"description":"Range of destination port for layer 4 protocol","type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}
		SRCPorts: InterfaceToPortType(data["src_ports"]),

		//{"description":"Range of source port for layer 4 protocol","type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}
		ProtocolID: data["protocol_id"].(int),

		//{"description":"Layer 4 protocol id in ip packet","type":"integer"}

	}
}

// InterfaceToFirewallServiceTypeSlice makes a slice of FirewallServiceType from interface
func InterfaceToFirewallServiceTypeSlice(data interface{}) []*FirewallServiceType {
	list := data.([]interface{})
	result := MakeFirewallServiceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceType(item))
	}
	return result
}

// MakeFirewallServiceTypeSlice() makes a slice of FirewallServiceType
func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
	return []*FirewallServiceType{}
}
