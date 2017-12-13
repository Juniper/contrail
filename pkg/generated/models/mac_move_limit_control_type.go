package models

// MACMoveLimitControlType

import "encoding/json"

// MACMoveLimitControlType
type MACMoveLimitControlType struct {
	MacMoveTimeWindow  MACMoveTimeWindow        `json:"mac_move_time_window"`
	MacMoveLimit       int                      `json:"mac_move_limit"`
	MacMoveLimitAction MACLimitExceedActionType `json:"mac_move_limit_action"`
}

// String returns json representation of the object
func (model *MACMoveLimitControlType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMACMoveLimitControlType makes MACMoveLimitControlType
func MakeMACMoveLimitControlType() *MACMoveLimitControlType {
	return &MACMoveLimitControlType{
		//TODO(nati): Apply default
		MacMoveLimit:       0,
		MacMoveLimitAction: MakeMACLimitExceedActionType(),
		MacMoveTimeWindow:  MakeMACMoveTimeWindow(),
	}
}

// InterfaceToMACMoveLimitControlType makes MACMoveLimitControlType from interface
func InterfaceToMACMoveLimitControlType(iData interface{}) *MACMoveLimitControlType {
	data := iData.(map[string]interface{})
	return &MACMoveLimitControlType{
		MacMoveTimeWindow: InterfaceToMACMoveTimeWindow(data["mac_move_time_window"]),

		//{"description":"MAC move time window","type":"integer","minimum":1,"maximum":60}
		MacMoveLimit: data["mac_move_limit"].(int),

		//{"description":"Number of MAC moves permitted in mac move time window","type":"integer"}
		MacMoveLimitAction: InterfaceToMACLimitExceedActionType(data["mac_move_limit_action"]),

		//{"description":"Action to be taken when MAC move limit exceeds","type":"string","enum":["log","alarm","shutdown","drop"]}

	}
}

// InterfaceToMACMoveLimitControlTypeSlice makes a slice of MACMoveLimitControlType from interface
func InterfaceToMACMoveLimitControlTypeSlice(data interface{}) []*MACMoveLimitControlType {
	list := data.([]interface{})
	result := MakeMACMoveLimitControlTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToMACMoveLimitControlType(item))
	}
	return result
}

// MakeMACMoveLimitControlTypeSlice() makes a slice of MACMoveLimitControlType
func MakeMACMoveLimitControlTypeSlice() []*MACMoveLimitControlType {
	return []*MACMoveLimitControlType{}
}
