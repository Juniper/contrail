package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeActionListType makes ActionListType
// nolint
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
// nolint
func InterfaceToActionListType(i interface{}) *ActionListType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ActionListType{
		//TODO(nati): Apply default
		GatewayName:           common.InterfaceToString(m["gateway_name"]),
		Log:                   common.InterfaceToBool(m["log"]),
		Alert:                 common.InterfaceToBool(m["alert"]),
		QosAction:             common.InterfaceToString(m["qos_action"]),
		AssignRoutingInstance: common.InterfaceToString(m["assign_routing_instance"]),
		MirrorTo:              InterfaceToMirrorActionType(m["mirror_to"]),
		SimpleAction:          common.InterfaceToString(m["simple_action"]),
		ApplyService:          common.InterfaceToStringList(m["apply_service"]),
	}
}

// MakeActionListTypeSlice() makes a slice of ActionListType
// nolint
func MakeActionListTypeSlice() []*ActionListType {
	return []*ActionListType{}
}

// InterfaceToActionListTypeSlice() makes a slice of ActionListType
// nolint
func InterfaceToActionListTypeSlice(i interface{}) []*ActionListType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ActionListType{}
	for _, item := range list {
		result = append(result, InterfaceToActionListType(item))
	}
	return result
}
