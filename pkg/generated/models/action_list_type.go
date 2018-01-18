package models

// ActionListType

import "encoding/json"

// ActionListType
type ActionListType struct {
	ApplyService          []string          `json:"apply_service,omitempty"`
	GatewayName           string            `json:"gateway_name,omitempty"`
	Log                   bool              `json:"log"`
	Alert                 bool              `json:"alert"`
	QosAction             string            `json:"qos_action,omitempty"`
	AssignRoutingInstance string            `json:"assign_routing_instance,omitempty"`
	MirrorTo              *MirrorActionType `json:"mirror_to,omitempty"`
	SimpleAction          SimpleActionType  `json:"simple_action,omitempty"`
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

// MakeActionListTypeSlice() makes a slice of ActionListType
func MakeActionListTypeSlice() []*ActionListType {
	return []*ActionListType{}
}
