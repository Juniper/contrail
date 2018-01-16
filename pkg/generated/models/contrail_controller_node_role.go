package models

// ContrailControllerNodeRole

import "encoding/json"

// ContrailControllerNodeRole
type ContrailControllerNodeRole struct {
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
}

// String returns json representation of the object
func (model *ContrailControllerNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeContrailControllerNodeRole makes ContrailControllerNodeRole
func MakeContrailControllerNodeRole() *ContrailControllerNodeRole {
	return &ContrailControllerNodeRole{
		//TODO(nati): Apply default
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ParentUUID:                "",
		ProvisioningLog:           "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ProvisioningState:         "",
		Annotations:               MakeKeyValuePairs(),
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
	}
}

// MakeContrailControllerNodeRoleSlice() makes a slice of ContrailControllerNodeRole
func MakeContrailControllerNodeRoleSlice() []*ContrailControllerNodeRole {
	return []*ContrailControllerNodeRole{}
}
