package models

// MACLimitControlType

// MACLimitControlType
//proteus:generate
type MACLimitControlType struct {
	MacLimit       int                      `json:"mac_limit,omitempty"`
	MacLimitAction MACLimitExceedActionType `json:"mac_limit_action,omitempty"`
}

// MakeMACLimitControlType makes MACLimitControlType
func MakeMACLimitControlType() *MACLimitControlType {
	return &MACLimitControlType{
		//TODO(nati): Apply default
		MacLimit:       0,
		MacLimitAction: MakeMACLimitExceedActionType(),
	}
}

// MakeMACLimitControlTypeSlice() makes a slice of MACLimitControlType
func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
	return []*MACLimitControlType{}
}
