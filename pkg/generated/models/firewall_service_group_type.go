package models

// FirewallServiceGroupType

import "encoding/json"

// FirewallServiceGroupType
type FirewallServiceGroupType struct {
	FirewallService []*FirewallServiceType `json:"firewall_service"`
}

//  parents relation object

// String returns json representation of the object
func (model *FirewallServiceGroupType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallServiceGroupType makes FirewallServiceGroupType
func MakeFirewallServiceGroupType() *FirewallServiceGroupType {
	return &FirewallServiceGroupType{
		//TODO(nati): Apply default

		FirewallService: MakeFirewallServiceTypeSlice(),
	}
}

// InterfaceToFirewallServiceGroupType makes FirewallServiceGroupType from interface
func InterfaceToFirewallServiceGroupType(iData interface{}) *FirewallServiceGroupType {
	data := iData.(map[string]interface{})
	return &FirewallServiceGroupType{

		FirewallService: InterfaceToFirewallServiceTypeSlice(data["firewall_service"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"dst_ports":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType","GoPremitive":false},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"DSTPorts","GoType":"PortType","GoPremitive":false},"protocol":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Protocol","GoType":"string","GoPremitive":true},"protocol_id":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProtocolID","GoType":"int","GoPremitive":true},"src_ports":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"EndPort","GoType":"L4PortType","GoPremitive":false},"start_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":-1,"Maximum":65535,"Ref":"types.json#/definitions/L4PortType","CollectionType":"","Column":"","Item":null,"GoName":"StartPort","GoType":"L4PortType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PortType","CollectionType":"","Column":"","Item":null,"GoName":"SRCPorts","GoType":"PortType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/FirewallServiceType","CollectionType":"","Column":"","Item":null,"GoName":"FirewallService","GoType":"FirewallServiceType","GoPremitive":false},"GoName":"FirewallService","GoType":"[]*FirewallServiceType","GoPremitive":true}

	}
}

// InterfaceToFirewallServiceGroupTypeSlice makes a slice of FirewallServiceGroupType from interface
func InterfaceToFirewallServiceGroupTypeSlice(data interface{}) []*FirewallServiceGroupType {
	list := data.([]interface{})
	result := MakeFirewallServiceGroupTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallServiceGroupType(item))
	}
	return result
}

// MakeFirewallServiceGroupTypeSlice() makes a slice of FirewallServiceGroupType
func MakeFirewallServiceGroupTypeSlice() []*FirewallServiceGroupType {
	return []*FirewallServiceGroupType{}
}
