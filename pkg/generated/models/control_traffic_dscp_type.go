package models

// ControlTrafficDscpType

import "encoding/json"

// ControlTrafficDscpType
type ControlTrafficDscpType struct {
	Control   DscpValueType `json:"control,omitempty"`
	Analytics DscpValueType `json:"analytics,omitempty"`
	DNS       DscpValueType `json:"dns,omitempty"`
}

// String returns json representation of the object
func (model *ControlTrafficDscpType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeControlTrafficDscpType makes ControlTrafficDscpType
func MakeControlTrafficDscpType() *ControlTrafficDscpType {
	return &ControlTrafficDscpType{
		//TODO(nati): Apply default
		Analytics: MakeDscpValueType(),
		DNS:       MakeDscpValueType(),
		Control:   MakeDscpValueType(),
	}
}

// MakeControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
func MakeControlTrafficDscpTypeSlice() []*ControlTrafficDscpType {
	return []*ControlTrafficDscpType{}
}
