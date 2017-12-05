package models

// FirewallServiceGroupType

import "encoding/json"

type FirewallServiceGroupType struct {
	FirewallService []*FirewallServiceType `json:"firewall_service"`
}

func (model *FirewallServiceGroupType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFirewallServiceGroupType() *FirewallServiceGroupType {
	return &FirewallServiceGroupType{
		//TODO(nati): Apply default

		FirewallService: MakeFirewallServiceTypeSlice(),
	}
}

func InterfaceToFirewallServiceGroupType(iData interface{}) *FirewallServiceGroupType {
	data := iData.(map[string]interface{})
	return &FirewallServiceGroupType{

		FirewallService: InterfaceToFirewallServiceTypeSlice(data["firewall_service"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dst_ports":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPorts","GoType":"PortType"},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string"},"protocol_id":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolID","GoType":"int"},"src_ports":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType"},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPorts","GoType":"PortType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallServiceType","CollectionType":"","Column":"","Item":null,"GoName":"FirewallService","GoType":"FirewallServiceType"},"GoName":"FirewallService","GoType":"[]*FirewallServiceType"}

	}
}

func InterfaceToFirewallServiceGroupTypeSlice(data interface{}) []*FirewallServiceGroupType {
	list := data.([]interface{})
	result := MakeFirewallServiceGroupTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceGroupType(item))
	}
	return result
}

func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
	return []*FirewallServiceGroupType{}
}
