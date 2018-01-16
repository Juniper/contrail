package models

// ContrailAnalyticsNode

import "encoding/json"

// ContrailAnalyticsNode
type ContrailAnalyticsNode struct {
	DisplayName               string         `json:"display_name,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
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
		ParentType:                "",
		Annotations:               MakeKeyValuePairs(),
		ProvisioningLog:           "",
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		UUID:                      "",
		ParentUUID:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		ProvisioningProgress:      0,
	}
}

// MakeContrailAnalyticsNodeSlice() makes a slice of ContrailAnalyticsNode
func MakeContrailAnalyticsNodeSlice() []*ContrailAnalyticsNode {
	return []*ContrailAnalyticsNode{}
}
