package models

// ControlTrafficDscpType

import "encoding/json"

// ControlTrafficDscpType
type ControlTrafficDscpType struct {
	Analytics DscpValueType `json:"analytics"`
	DNS       DscpValueType `json:"dns"`
	Control   DscpValueType `json:"control"`
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
		Control:   MakeDscpValueType(),
		Analytics: MakeDscpValueType(),
		DNS:       MakeDscpValueType(),
	}
}

// MakeControlTrafficDscpTypeSlice() makes a slice of ControlTrafficDscpType
func MakeControlTrafficDscpTypeSlice() []*ControlTrafficDscpType {
	return []*ControlTrafficDscpType{}
}
