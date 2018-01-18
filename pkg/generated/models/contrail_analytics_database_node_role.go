package models

// ContrailAnalyticsDatabaseNodeRole

import "encoding/json"

// ContrailAnalyticsDatabaseNodeRole
type ContrailAnalyticsDatabaseNodeRole struct {
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
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
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ProvisioningLog:           "",
		ProvisioningProgressStage: "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		Perms2:                MakePermType2(),
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
		ProvisioningProgress:  0,
		ProvisioningStartTime: "",
		ProvisioningState:     "",
	}
}

// MakeContrailAnalyticsDatabaseNodeRoleSlice() makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
	return []*ContrailAnalyticsDatabaseNodeRole{}
}
