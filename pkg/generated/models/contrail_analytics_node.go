package models

// ContrailAnalyticsNode

import "encoding/json"

// ContrailAnalyticsNode
type ContrailAnalyticsNode struct {
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
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
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		ProvisioningState:     "",
		ProvisioningLog:       "",
		ProvisioningProgress:  0,
		ProvisioningStartTime: "",
		IDPerms:               MakeIdPermsType(),
		ParentUUID:            "",
		ParentType:            "",
		FQName:                []string{},
		ProvisioningProgressStage: "",
		UUID: "",
	}
}

// MakeContrailAnalyticsNodeSlice() makes a slice of ContrailAnalyticsNode
func MakeContrailAnalyticsNodeSlice() []*ContrailAnalyticsNode {
	return []*ContrailAnalyticsNode{}
}
