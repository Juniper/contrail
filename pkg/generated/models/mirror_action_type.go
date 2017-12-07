package models

// MirrorActionType

import "encoding/json"

// MirrorActionType
type MirrorActionType struct {
	Encapsulation            string              `json:"encapsulation"`
	AnalyzerMacAddress       string              `json:"analyzer_mac_address"`
	NicAssistedMirroringVlan VlanIdType          `json:"nic_assisted_mirroring_vlan"`
	NHMode                   NHModeType          `json:"nh_mode"`
	UDPPort                  int                 `json:"udp_port"`
	RoutingInstance          string              `json:"routing_instance"`
	AnalyzerIPAddress        string              `json:"analyzer_ip_address"`
	AnalyzerName             string              `json:"analyzer_name"`
	JuniperHeader            bool                `json:"juniper_header"`
	StaticNHHeader           *StaticMirrorNhType `json:"static_nh_header"`
	NicAssistedMirroring     bool                `json:"nic_assisted_mirroring"`
}

//  parents relation object

// String returns json representation of the object
func (model *MirrorActionType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMirrorActionType makes MirrorActionType
func MakeMirrorActionType() *MirrorActionType {
	return &MirrorActionType{
		//TODO(nati): Apply default
		StaticNHHeader:           MakeStaticMirrorNhType(),
		NicAssistedMirroring:     false,
		AnalyzerName:             "",
		JuniperHeader:            false,
		UDPPort:                  0,
		RoutingInstance:          "",
		AnalyzerIPAddress:        "",
		Encapsulation:            "",
		AnalyzerMacAddress:       "",
		NicAssistedMirroringVlan: MakeVlanIdType(),
		NHMode: MakeNHModeType(),
	}
}

// InterfaceToMirrorActionType makes MirrorActionType from interface
func InterfaceToMirrorActionType(iData interface{}) *MirrorActionType {
	data := iData.(map[string]interface{})
	return &MirrorActionType{
		NicAssistedMirroringVlan: InterfaceToVlanIdType(data["nic_assisted_mirroring_vlan"]),

		//{"Title":"","Description":"The VLAN to be tagged on the traffic for NIC to Mirror","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType","GoPremitive":false}
		NHMode: InterfaceToNHModeType(data["nh_mode"]),

		//{"Title":"","Description":"This mode used to determine static or dynamic nh ","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType","GoPremitive":false}
		UDPPort: data["udp_port"].(int),

		//{"Title":"","Description":"ip udp port used in contrail default encapsulation for mirroring","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int","GoPremitive":true}
		RoutingInstance: data["routing_instance"].(string),

		//{"Title":"","Description":"Internal use only, should be set to -1","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string","GoPremitive":true}
		AnalyzerIPAddress: data["analyzer_ip_address"].(string),

		//{"Title":"","Description":"ip address of interface to which mirrored packets are sent","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string","GoPremitive":true}
		Encapsulation: data["encapsulation"].(string),

		//{"Title":"","Description":"Encapsulation for Mirrored packet, not used currently","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string","GoPremitive":true}
		AnalyzerMacAddress: data["analyzer_mac_address"].(string),

		//{"Title":"","Description":"mac address of interface to which mirrored packets are sent ","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string","GoPremitive":true}
		AnalyzerName: data["analyzer_name"].(string),

		//{"Title":"","Description":"Name of service instance used as analyzer","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string","GoPremitive":true}
		JuniperHeader: data["juniper_header"].(bool),

		//{"Title":"","Description":"This flag is used to determine with/without juniper-header","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool","GoPremitive":true}
		StaticNHHeader: InterfaceToStaticMirrorNhType(data["static_nh_header"]),

		//{"Title":"","Description":"vtep details required if static nh enabled","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType","GoPremitive":false},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string","GoPremitive":true},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType","GoPremitive":false}
		NicAssistedMirroring: data["nic_assisted_mirroring"].(bool),

		//{"Title":"","Description":"This flag is used to select nic assisted mirroring","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool","GoPremitive":true}

	}
}

// InterfaceToMirrorActionTypeSlice makes a slice of MirrorActionType from interface
func InterfaceToMirrorActionTypeSlice(data interface{}) []*MirrorActionType {
	list := data.([]interface{})
	result := MakeMirrorActionTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMirrorActionType(item))
	}
	return result
}

// MakeMirrorActionTypeSlice() makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
	return []*MirrorActionType{}
}
