package models

// VrfAssignTableType

// VrfAssignTableType
//proteus:generate
type VrfAssignTableType struct {
	VRFAssignRule []*VrfAssignRuleType `json:"vrf_assign_rule,omitempty"`
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
