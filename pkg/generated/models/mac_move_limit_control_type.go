package models

// MACMoveLimitControlType

// MACMoveLimitControlType
//proteus:generate
type MACMoveLimitControlType struct {
	MacMoveTimeWindow  MACMoveTimeWindow        `json:"mac_move_time_window,omitempty"`
	MacMoveLimit       int                      `json:"mac_move_limit,omitempty"`
	MacMoveLimitAction MACLimitExceedActionType `json:"mac_move_limit_action,omitempty"`
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
