package models

// VrfAssignRuleType

import "encoding/json"

// VrfAssignRuleType
type VrfAssignRuleType struct {
	RoutingInstance string              `json:"routing_instance,omitempty"`
	MatchCondition  *MatchConditionType `json:"match_condition,omitempty"`
	VlanTag         int                 `json:"vlan_tag,omitempty"`
	IgnoreACL       bool                `json:"ignore_acl,omitempty"`
}

// String returns json representation of the object
func (model *VrfAssignRuleType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVrfAssignRuleType makes VrfAssignRuleType
func MakeVrfAssignRuleType() *VrfAssignRuleType {
	return &VrfAssignRuleType{
		//TODO(nati): Apply default
		VlanTag:         0,
		IgnoreACL:       false,
		RoutingInstance: "",
		MatchCondition:  MakeMatchConditionType(),
	}
}

// MakeVrfAssignRuleTypeSlice() makes a slice of VrfAssignRuleType
func MakeVrfAssignRuleTypeSlice() []*VrfAssignRuleType {
	return []*VrfAssignRuleType{}
}
