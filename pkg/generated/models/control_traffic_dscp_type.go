package models

// ControlTrafficDscpType

// ControlTrafficDscpType
//proteus:generate
type ControlTrafficDscpType struct {
	Control   DscpValueType `json:"control,omitempty"`
	Analytics DscpValueType `json:"analytics,omitempty"`
	DNS       DscpValueType `json:"dns,omitempty"`
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
