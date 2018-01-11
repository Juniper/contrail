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
		MacMoveTimeWindow:  MakeMACMoveTimeWindow(),
		MacMoveLimit:       0,
		MacMoveLimitAction: MakeMACLimitExceedActionType(),
	}
}

// MakeMACMoveLimitControlTypeSlice() makes a slice of MACMoveLimitControlType
func MakeMACMoveLimitControlTypeSlice() []*MACMoveLimitControlType {
	return []*MACMoveLimitControlType{}
}
