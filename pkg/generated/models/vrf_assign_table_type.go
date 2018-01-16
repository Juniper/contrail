package models

// VrfAssignTableType

import "encoding/json"

// VrfAssignTableType
type VrfAssignTableType struct {
	VRFAssignRule []*VrfAssignRuleType `json:"vrf_assign_rule,omitempty"`
}

// String returns json representation of the object
func (model *VrfAssignTableType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVrfAssignTableType makes VrfAssignTableType
func MakeVrfAssignTableType() *VrfAssignTableType {
	return &VrfAssignTableType{
		//TODO(nati): Apply default

		VRFAssignRule: MakeVrfAssignRuleTypeSlice(),
	}
}

// MakeVrfAssignTableTypeSlice() makes a slice of VrfAssignTableType
func MakeVrfAssignTableTypeSlice() []*VrfAssignTableType {
	return []*VrfAssignTableType{}
}
