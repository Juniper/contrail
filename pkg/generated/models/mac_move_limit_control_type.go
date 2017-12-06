package models

// MACMoveLimitControlType

import "encoding/json"

// MACMoveLimitControlType
type MACMoveLimitControlType struct {
	MacMoveTimeWindow  MACMoveTimeWindow        `json:"mac_move_time_window"`
	MacMoveLimit       int                      `json:"mac_move_limit"`
	MacMoveLimitAction MACLimitExceedActionType `json:"mac_move_limit_action"`
}

//  parents relation object

// String returns json representation of the object
func (model *MACMoveLimitControlType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeMACMoveLimitControlType makes MACMoveLimitControlType
func MakeMACMoveLimitControlType() *MACMoveLimitControlType {
	return &MACMoveLimitControlType{
		//TODO(nati): Apply default
		MacMoveTimeWindow:  MakeMACMoveTimeWindow(),
		MacMoveLimit:       0,
		MacMoveLimitAction: MakeMACLimitExceedActionType(),
	}
}

// InterfaceToMACMoveLimitControlType makes MACMoveLimitControlType from interface
func InterfaceToMACMoveLimitControlType(iData interface{}) *MACMoveLimitControlType {
	data := iData.(map[string]interface{})
	return &MACMoveLimitControlType{
		MacMoveTimeWindow: InterfaceToMACMoveTimeWindow(data["mac_move_time_window"]),

		//{"Title":"","Description":"MAC move time window","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":60,"Ref":"types.json#/definitions/MACMoveTimeWindow","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveTimeWindow","GoType":"MACMoveTimeWindow","GoPremitive":false}
		MacMoveLimit: data["mac_move_limit"].(int),

		//{"Title":"","Description":"Number of MAC moves permitted in mac move time window","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveLimit","GoType":"int","GoPremitive":true}
		MacMoveLimitAction: InterfaceToMACLimitExceedActionType(data["mac_move_limit_action"]),

		//{"Title":"","Description":"Action to be taken when MAC move limit exceeds","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["log","alarm","shutdown","drop"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/MACLimitExceedActionType","CollectionType":"","Column":"","Item":null,"GoName":"MacMoveLimitAction","GoType":"MACLimitExceedActionType","GoPremitive":false}

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
