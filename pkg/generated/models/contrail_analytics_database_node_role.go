package models

// ContrailAnalyticsDatabaseNodeRole

import "encoding/json"

// ContrailAnalyticsDatabaseNodeRole
type ContrailAnalyticsDatabaseNodeRole struct {
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
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
		ProvisioningState:         "",
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		ProvisioningProgressStage: "",
		Perms2:                MakePermType2(),
		UUID:                  "",
		ParentUUID:            "",
		ProvisioningStartTime: "",
		ProvisioningLog:       "",
		ProvisioningProgress:  0,
	}
}

// MakeContrailAnalyticsDatabaseNodeRoleSlice() makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
	return []*ContrailAnalyticsDatabaseNodeRole{}
}
