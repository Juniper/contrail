package models

// ContrailAnalyticsNode

import "encoding/json"

// ContrailAnalyticsNode
type ContrailAnalyticsNode struct {
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
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
		ProvisioningStartTime:     "",
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ProvisioningProgressStage: "",
		FQName:               []string{},
		Annotations:          MakeKeyValuePairs(),
		ProvisioningState:    "",
		ProvisioningLog:      "",
		ProvisioningProgress: 0,
	}
}

// MakeContrailAnalyticsNodeSlice() makes a slice of ContrailAnalyticsNode
func MakeContrailAnalyticsNodeSlice() []*ContrailAnalyticsNode {
	return []*ContrailAnalyticsNode{}
}
