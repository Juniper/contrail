package models

// ContrailAnalyticsNode

import "encoding/json"

// ContrailAnalyticsNode
type ContrailAnalyticsNode struct {
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
}

// String returns json representation of the object
func (model *ContrailAnalyticsNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeContrailAnalyticsNode makes ContrailAnalyticsNode
func MakeContrailAnalyticsNode() *ContrailAnalyticsNode {
	return &ContrailAnalyticsNode{
		//TODO(nati): Apply default
		Perms2:                    MakePermType2(),
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		ProvisioningProgressStage: "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentUUID:                "",
		FQName:                    []string{},
		ProvisioningState:         "",
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningStartTime:     "",
		DisplayName:               "",
	}
}

// MakeContrailAnalyticsNodeSlice() makes a slice of ContrailAnalyticsNode
func MakeContrailAnalyticsNodeSlice() []*ContrailAnalyticsNode {
	return []*ContrailAnalyticsNode{}
}
