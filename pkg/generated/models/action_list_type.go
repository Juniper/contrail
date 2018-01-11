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

// String returns json representation of the object
func (model *ActionListType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeActionListType makes ActionListType
func MakeActionListType() *ActionListType {
	return &ActionListType{
		//TODO(nati): Apply default
		Alert:                 false,
		QosAction:             "",
		AssignRoutingInstance: "",
		MirrorTo:              MakeMirrorActionType(),
		SimpleAction:          MakeSimpleActionType(),
		ApplyService:          []string{},
		GatewayName:           "",
		Log:                   false,
	}
}

// MakeActionListTypeSlice() makes a slice of ActionListType
func MakeActionListTypeSlice() []*ActionListType {
	return []*ActionListType{}
}
