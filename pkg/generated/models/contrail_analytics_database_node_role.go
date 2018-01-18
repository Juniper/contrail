package models

// ContrailAnalyticsDatabaseNodeRole

import "encoding/json"

// ContrailAnalyticsDatabaseNodeRole
type ContrailAnalyticsDatabaseNodeRole struct {
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
}

// String returns json representation of the object
func (model *ContrailAnalyticsDatabaseNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeContrailAnalyticsDatabaseNodeRole makes ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRole() *ContrailAnalyticsDatabaseNodeRole {
	return &ContrailAnalyticsDatabaseNodeRole{
		//TODO(nati): Apply default
		IDPerms:                   MakeIdPermsType(),
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ParentUUID:                "",
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		FQName:                    []string{},
		ProvisioningProgress:      0,
		Annotations:               MakeKeyValuePairs(),
		ParentType:                "",
		ProvisioningLog:           "",
		DisplayName:               "",
	}
}

// MakeContrailAnalyticsDatabaseNodeRoleSlice() makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
	return []*ContrailAnalyticsDatabaseNodeRole{}
}
