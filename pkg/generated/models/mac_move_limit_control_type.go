package models

// MACMoveLimitControlType

import "encoding/json"

type MACMoveLimitControlType struct {
	MacMoveTimeWindow  MACMoveTimeWindow        `json:"mac_move_time_window"`
	MacMoveLimit       int                      `json:"mac_move_limit"`
	MacMoveLimitAction MACLimitExceedActionType `json:"mac_move_limit_action"`
}

func (model *MACMoveLimitControlType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeMACMoveLimitControlType() *MACMoveLimitControlType {
	return &MACMoveLimitControlType{
		//TODO(nati): Apply default
		MacMoveLimit:       0,
		MacMoveLimitAction: MakeMACLimitExceedActionType(),
		MacMoveTimeWindow:  MakeMACMoveTimeWindow(),
	}
}

func InterfaceToMACMoveLimitControlType(iData interface{}) *MACMoveLimitControlType {
	data := iData.(map[string]interface{})
	return &MACMoveLimitControlType{
		MacMoveLimit: data["mac_move_limit"].(int),

		//{"Title":"","Description":"Number of MAC moves permitted in mac move time window","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveLimit","GoType":"int"}
		MacMoveLimitAction: InterfaceToMACLimitExceedActionType(data["mac_move_limit_action"]),

		//{"Title":"","Description":"Action to be taken when MAC move limit exceeds","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveLimitAction","GoType":"MACLimitExceedActionType"}
		MacMoveTimeWindow: InterfaceToMACMoveTimeWindow(data["mac_move_time_window"]),

		//{"Title":"","Description":"MAC move time window","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":60,"Ref":"types.json#/definitions/MACMoveTimeWindow","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveTimeWindow","GoType":"MACMoveTimeWindow"}

	}
}

func InterfaceToMACMoveLimitControlTypeSlice(data interface{}) []*MACMoveLimitControlType {
	list := data.([]interface{})
	result := MakeMACMoveLimitControlTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACMoveLimitControlType(item))
	}
	return result
}

func MakeMACMoveLimitControlTypeSlice() []*MACMoveLimitControlType {
	return []*MACMoveLimitControlType{}
}
