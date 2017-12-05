package models

// ActionListType

import "encoding/json"

type ActionListType struct {
	AssignRoutingInstance string            `json:"assign_routing_instance"`
	MirrorTo              *MirrorActionType `json:"mirror_to"`
	SimpleAction          SimpleActionType  `json:"simple_action"`
	ApplyService          []string          `json:"apply_service"`
	GatewayName           string            `json:"gateway_name"`
	Log                   bool              `json:"log"`
	Alert                 bool              `json:"alert"`
	QosAction             string            `json:"qos_action"`
}

func (model *ActionListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeActionListType() *ActionListType {
	return &ActionListType{
		//TODO(nati): Apply default
		MirrorTo:              MakeMirrorActionType(),
		SimpleAction:          MakeSimpleActionType(),
		ApplyService:          []string{},
		GatewayName:           "",
		Log:                   false,
		Alert:                 false,
		QosAction:             "",
		AssignRoutingInstance: "",
	}
}

func InterfaceToActionListType(iData interface{}) *ActionListType {
	data := iData.(map[string]interface{})
	return &ActionListType{
		GatewayName: data["gateway_name"].(string),

		//{"Title":"","Description":"For internal use only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"GatewayName","GoType":"string"}
		Log: data["log"].(bool),

		//{"Title":"","Description":"Flow records for traffic matching this rule are sent at higher priority","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Log","GoType":"bool"}
		Alert: data["alert"].(bool),

		//{"Title":"","Description":"For internal use only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Alert","GoType":"bool"}
		QosAction: data["qos_action"].(string),

		//{"Title":"","Description":"FQN of Qos configuration object for QoS marking","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"QosAction","GoType":"string"}
		AssignRoutingInstance: data["assign_routing_instance"].(string),

		//{"Title":"","Description":"For internal use only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AssignRoutingInstance","GoType":"string"}
		MirrorTo: InterfaceToMirrorActionType(data["mirror_to"]),

		//{"Title":"","Description":"Mirror traffic matching this rule","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"analyzer_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerIPAddress","GoType":"string"},"analyzer_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerMacAddress","GoType":"string"},"analyzer_name":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AnalyzerName","GoType":"string"},"encapsulation":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Encapsulation","GoType":"string"},"juniper_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"JuniperHeader","GoType":"bool"},"nh_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["dynamic","static"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/NHModeType","CollectionType":"","Column":"","Item":null,"GoName":"NHMode","GoType":"NHModeType"},"nic_assisted_mirroring":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroring","GoType":"bool"},"nic_assisted_mirroring_vlan":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"NicAssistedMirroringVlan","GoType":"VlanIdType"},"routing_instance":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RoutingInstance","GoType":"string"},"static_nh_header":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"vni":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":16777215,"Ref":"types.json#/definitions/VxlanNetworkIdentifierType","CollectionType":"","Column":"","Item":null,"GoName":"Vni","GoType":"VxlanNetworkIdentifierType"},"vtep_dst_ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTIPAddress","GoType":"string"},"vtep_dst_mac_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VtepDSTMacAddress","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/StaticMirrorNhType","CollectionType":"","Column":"","Item":null,"GoName":"StaticNHHeader","GoType":"StaticMirrorNhType"},"udp_port":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"UDPPort","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MirrorActionType","CollectionType":"","Column":"","Item":null,"GoName":"MirrorTo","GoType":"MirrorActionType"}
		SimpleAction: InterfaceToSimpleActionType(data["simple_action"]),

		//{"Title":"","Description":"Simple allow(pass) or deny action for traffic matching this rule","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["deny","pass"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SimpleActionType","CollectionType":"","Column":"","Item":null,"GoName":"SimpleAction","GoType":"SimpleActionType"}
		ApplyService: data["apply_service"].([]string),

		//{"Title":"","Description":"Ordered list of service instances in service chain applied to traffic matching the rule","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ApplyService","GoType":"string"},"GoName":"ApplyService","GoType":"[]string"}

	}
}

func InterfaceToActionListTypeSlice(data interface{}) []*ActionListType {
	list := data.([]interface{})
	result := MakeActionListTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToActionListType(item))
	}
	return result
}

func MakeActionListTypeSlice() []*ActionListType {
	return []*ActionListType{}
}
