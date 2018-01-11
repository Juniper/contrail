package models

// ContrailAnalyticsDatabaseNodeRole

import "encoding/json"

// ContrailAnalyticsDatabaseNodeRole
type ContrailAnalyticsDatabaseNodeRole struct {
	Perms2                    *PermType2     `json:"perms2"`
	ParentType                string         `json:"parent_type"`
	ProvisioningLog           string         `json:"provisioning_log"`
	ProvisioningProgress      int            `json:"provisioning_progress"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage"`
	Annotations               *KeyValuePairs `json:"annotations"`
	UUID                      string         `json:"uuid"`
	ParentUUID                string         `json:"parent_uuid"`
	FQName                    []string       `json:"fq_name"`
	IDPerms                   *IdPermsType   `json:"id_perms"`
	DisplayName               string         `json:"display_name"`
	ProvisioningStartTime     string         `json:"provisioning_start_time"`
	ProvisioningState         string         `json:"provisioning_state"`
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
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		Perms2:                MakePermType2(),
		ParentType:            "",
		ParentUUID:            "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		ProvisioningStartTime: "",
		ProvisioningState:     "",
		Annotations:           MakeKeyValuePairs(),
		UUID:                  "",
	}
}

// MakeContrailAnalyticsDatabaseNodeRoleSlice() makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
	return []*ContrailAnalyticsDatabaseNodeRole{}
}
