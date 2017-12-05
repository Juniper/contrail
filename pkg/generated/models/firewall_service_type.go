package models

// FirewallServiceType

import "encoding/json"

type FirewallServiceType struct {
	Protocol   string    `json:"protocol"`
	DSTPorts   *PortType `json:"dst_ports"`
	SRCPorts   *PortType `json:"src_ports"`
	ProtocolID int       `json:"protocol_id"`
}

func (model *FirewallServiceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFirewallServiceType() *FirewallServiceType {
	return &FirewallServiceType{
		//TODO(nati): Apply default
		Protocol:   "",
		DSTPorts:   MakePortType(),
		SRCPorts:   MakePortType(),
		ProtocolID: 0,
	}
}

func InterfaceToFirewallServiceType(iData interface{}) *FirewallServiceType {
	data := iData.(map[string]interface{})
	return &FirewallServiceType{
		Protocol: data["protocol"].(string),

		//{"Title":"","Description":"Layer 4 protocol in ip packet","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"}
		DSTPorts: InterfaceToPortType(data["dst_ports"]),

		//{"Title":"","Description":"Range of destination port for layer 4 protocol","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPorts","GoType":"PortType"}
		SRCPorts: InterfaceToPortType(data["src_ports"]),

		//{"Title":"","Description":"Range of source port for layer 4 protocol","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPorts","GoType":"PortType"}
		ProtocolID: data["protocol_id"].(int),

		//{"Title":"","Description":"Layer 4 protocol id in ip packet","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolID","GoType":"int"}

	}
}

func InterfaceToFirewallServiceTypeSlice(data interface{}) []*FirewallServiceType {
	list := data.([]interface{})
	result := MakeFirewallServiceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceType(item))
	}
	return result
}

func MakeFirewallServiceTypeSlice() []*FirewallServiceType {
	return []*FirewallServiceType{}
}
