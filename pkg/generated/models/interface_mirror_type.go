package models

// InterfaceMirrorType

import "encoding/json"

// InterfaceMirrorType
type InterfaceMirrorType struct {
	MirrorTo         *MirrorActionType    `json:"mirror_to"`
	TrafficDirection TrafficDirectionType `json:"traffic_direction"`
}

//  parents relation object

// String returns json representation of the object
func (model *InterfaceMirrorType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeInterfaceMirrorType makes InterfaceMirrorType
func MakeInterfaceMirrorType() *InterfaceMirrorType {
	return &InterfaceMirrorType{
		//TODO(nati): Apply default
		TrafficDirection: MakeTrafficDirectionType(),
		MirrorTo:         MakeMirrorActionType(),
	}
}

// InterfaceToInterfaceMirrorType makes InterfaceMirrorType from interface
func InterfaceToInterfaceMirrorType(iData interface{}) *InterfaceMirrorType {
	data := iData.(map[string]interface{})
	return &InterfaceMirrorType{
		TrafficDirection: InterfaceToTrafficDirectionType(data["traffic_direction"]),

		//{"Title":"","Description":"Specifies direction of traffic to mirror, Ingress, Egress or both","SQL":"","Default":"both","Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["ingress","egress","both"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/TrafficDirectionType","CollectionType":"","Column":"","Item":null,"GoName":"TrafficDirection","GoType":"TrafficDirectionType","GoPremitive":false}
		MirrorTo: InterfaceToMirrorActionType(data["mirror_to"]),

		//{"Title":"","Description":"Mirror destination configuration","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string","GoPremitive":true},"analyzer_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string","GoPremitive":true},"analyzer_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string","GoPremitive":true},"encapsulation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string","GoPremitive":true},"juniper_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool","GoPremitive":true},"nh_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType","GoPremitive":false},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool","GoPremitive":true},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType","GoPremitive":false},"routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string","GoPremitive":true},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType","GoPremitive":false},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string","GoPremitive":true},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType","GoPremitive":false},"udp_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType","GoPremitive":false}

	}
}

// InterfaceToInterfaceMirrorTypeSlice makes a slice of InterfaceMirrorType from interface
func InterfaceToInterfaceMirrorTypeSlice(data interface{}) []*InterfaceMirrorType {
	list := data.([]interface{})
	result := MakeInterfaceMirrorTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToInterfaceMirrorType(item))
	}
	return result
}

// MakeInterfaceMirrorTypeSlice() makes a slice of InterfaceMirrorType
func MakeInterfaceMirrorTypeSlice() []*InterfaceMirrorType {
	return []*InterfaceMirrorType{}
}
