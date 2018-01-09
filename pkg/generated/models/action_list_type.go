package models

// ActionListType

import "encoding/json"

// ActionListType
type ActionListType struct {
	MirrorTo              *MirrorActionType `json:"mirror_to"`
	SimpleAction          SimpleActionType  `json:"simple_action"`
	ApplyService          []string          `json:"apply_service"`
	GatewayName           string            `json:"gateway_name"`
	Log                   bool              `json:"log"`
	Alert                 bool              `json:"alert"`
	QosAction             string            `json:"qos_action"`
	AssignRoutingInstance string            `json:"assign_routing_instance"`
}

// String returns json representation of the object
func (model *ActionListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeActionListType makes ActionListType
func MakeActionListType() *ActionListType {
	return &ActionListType{
		//TODO(nati): Apply default
		QosAction:             "",
		AssignRoutingInstance: "",
		MirrorTo:              MakeMirrorActionType(),
		SimpleAction:          MakeSimpleActionType(),
		ApplyService:          []string{},
		GatewayName:           "",
		Log:                   false,
		Alert:                 false,
	}
}

// InterfaceToActionListType makes ActionListType from interface
func InterfaceToActionListType(iData interface{}) *ActionListType {
	data := iData.(map[string]interface{})
	return &ActionListType{
		GatewayName: data["gateway_name"].(string),

		//{"description":"For internal use only","type":"string"}
		Log: data["log"].(bool),

		//{"description":"Flow records for traffic matching this rule are sent at higher priority","type":"boolean"}
		Alert: data["alert"].(bool),

		//{"description":"For internal use only","type":"boolean"}
		QosAction: data["qos_action"].(string),

		//{"description":"FQN of Qos configuration object for QoS marking","type":"string"}
		AssignRoutingInstance: data["assign_routing_instance"].(string),

		//{"description":"For internal use only","type":"string"}
		MirrorTo: InterfaceToMirrorActionType(data["mirror_to"]),

		//{"description":"Mirror traffic matching this rule","type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}}
		SimpleAction: InterfaceToSimpleActionType(data["simple_action"]),

		//{"description":"Simple allow(pass) or deny action for traffic matching this rule","type":"string","enum":["deny","pass"]}
		ApplyService: data["apply_service"].([]string),

		//{"description":"Ordered list of service instances in service chain applied to traffic matching the rule","type":"array","item":{"type":"string"}}

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
