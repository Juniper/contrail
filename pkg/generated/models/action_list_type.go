package models

// ActionListType

import "encoding/json"

// ActionListType
type ActionListType struct {
	ApplyService          []string          `json:"apply_service"`
	GatewayName           string            `json:"gateway_name"`
	Log                   bool              `json:"log"`
	Alert                 bool              `json:"alert"`
	QosAction             string            `json:"qos_action"`
	AssignRoutingInstance string            `json:"assign_routing_instance"`
	MirrorTo              *MirrorActionType `json:"mirror_to"`
	SimpleAction          SimpleActionType  `json:"simple_action"`
}

//  parents relation object

// String returns json representation of the object
func (model *ActionListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeActionListType makes ActionListType
func MakeActionListType() *ActionListType {
	return &ActionListType{
		//TODO(nati): Apply default
		SimpleAction:          MakeSimpleActionType(),
		ApplyService:          []string{},
		GatewayName:           "",
		Log:                   false,
		Alert:                 false,
		QosAction:             "",
		AssignRoutingInstance: "",
		MirrorTo:              MakeMirrorActionType(),
	}
}

// InterfaceToActionListType makes ActionListType from interface
func InterfaceToActionListType(iData interface{}) *ActionListType {
	data := iData.(map[string]interface{})
	return &ActionListType{
		ApplyService: data["apply_service"].([]string),

		//{"Title":"","Description":"Ordered list of service instances in service chain applied to traffic matching the rule","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ApplyService","GoType":"string","GoPremitive":true},"GoName":"ApplyService","GoType":"[]string","GoPremitive":true}
		GatewayName: data["gateway_name"].(string),

		//{"Title":"","Description":"For internal use only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"GatewayName","GoType":"string","GoPremitive":true}
		Log: data["log"].(bool),

		//{"Title":"","Description":"Flow records for traffic matching this rule are sent at higher priority","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Log","GoType":"bool","GoPremitive":true}
		Alert: data["alert"].(bool),

		//{"Title":"","Description":"For internal use only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Alert","GoType":"bool","GoPremitive":true}
		QosAction: data["qos_action"].(string),

		//{"Title":"","Description":"FQN of Qos configuration object for QoS marking","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"QosAction","GoType":"string","GoPremitive":true}
		AssignRoutingInstance: data["assign_routing_instance"].(string),

		//{"Title":"","Description":"For internal use only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AssignRoutingInstance","GoType":"string","GoPremitive":true}
		MirrorTo: InterfaceToMirrorActionType(data["mirror_to"]),

		//{"Title":"","Description":"Mirror traffic matching this rule","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string","GoPremitive":true},"analyzer_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string","GoPremitive":true},"analyzer_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string","GoPremitive":true},"encapsulation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string","GoPremitive":true},"juniper_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool","GoPremitive":true},"nh_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType","GoPremitive":false},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool","GoPremitive":true},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType","GoPremitive":false},"routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string","GoPremitive":true},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType","GoPremitive":false},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string","GoPremitive":true},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType","GoPremitive":false},"udp_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType","GoPremitive":false}
		SimpleAction: InterfaceToSimpleActionType(data["simple_action"]),

		//{"Title":"","Description":"Simple allow(pass) or deny action for traffic matching this rule","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["deny","pass"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SimpleActionType","CollectionType":"","Column":"","Item":null,"GoName":"SimpleAction","GoType":"SimpleActionType","GoPremitive":false}

	}
}

// InterfaceToActionListTypeSlice makes a slice of ActionListType from interface
func InterfaceToActionListTypeSlice(data interface{}) []*ActionListType {
	list := data.([]interface{})
	result := MakeActionListTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToActionListType(item))
	}
	return result
}

// MakeActionListTypeSlice() makes a slice of ActionListType
func MakeActionListTypeSlice() []*ActionListType {
	return []*ActionListType{}
}
