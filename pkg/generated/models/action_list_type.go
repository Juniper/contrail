package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeActionListType makes ActionListType
func MakeActionListType() *ActionListType {
	return &ActionListType{
		//TODO(nati): Apply default
		GatewayName:           "",
		Log:                   false,
		Alert:                 false,
		QosAction:             "",
		AssignRoutingInstance: "",
		MirrorTo:              MakeMirrorActionType(),
		SimpleAction:          "",
		ApplyService:          []string{},
	}
}

// MakeActionListType makes ActionListType
func InterfaceToActionListType(i interface{}) *ActionListType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ActionListType{
		//TODO(nati): Apply default
		GatewayName:           schema.InterfaceToString(m["gateway_name"]),
		Log:                   schema.InterfaceToBool(m["log"]),
		Alert:                 schema.InterfaceToBool(m["alert"]),
		QosAction:             schema.InterfaceToString(m["qos_action"]),
		AssignRoutingInstance: schema.InterfaceToString(m["assign_routing_instance"]),
		MirrorTo:              InterfaceToMirrorActionType(m["mirror_to"]),
		SimpleAction:          schema.InterfaceToString(m["simple_action"]),
		ApplyService:          schema.InterfaceToStringList(m["apply_service"]),
	}
}

// MakeActionListTypeSlice() makes a slice of ActionListType
func MakeActionListTypeSlice() []*ActionListType {
	return []*ActionListType{}
}

// InterfaceToActionListTypeSlice() makes a slice of ActionListType
func InterfaceToActionListTypeSlice(i interface{}) []*ActionListType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ActionListType{}
	for _, item := range list {
		result = append(result, InterfaceToActionListType(item))
	}
	return result
}
