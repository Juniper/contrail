package models

// VrfAssignRuleType

// VrfAssignRuleType
//proteus:generate
type VrfAssignRuleType struct {
	RoutingInstance string              `json:"routing_instance,omitempty"`
	MatchCondition  *MatchConditionType `json:"match_condition,omitempty"`
	VlanTag         int                 `json:"vlan_tag,omitempty"`
	IgnoreACL       bool                `json:"ignore_acl"`
}

// MakeVrfAssignRuleType makes VrfAssignRuleType
func MakeVrfAssignRuleType() *VrfAssignRuleType {
	return &VrfAssignRuleType{
		//TODO(nati): Apply default
		RoutingInstance: "",
		MatchCondition:  MakeMatchConditionType(),
		VlanTag:         0,
		IgnoreACL:       false,
	}
}

// MakeVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
	return []*VrfAssignRuleType{}
}
