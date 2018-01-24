package models

// ContrailControllerNodeRole

import "encoding/json"

// ContrailControllerNodeRole
type ContrailControllerNodeRole struct {
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
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
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		ProvisioningLog:           "",
		IDPerms:                   MakeIdPermsType(),
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningProgress:      0,
		ProvisioningState:         "",
		Perms2:                    MakePermType2(),
	}
}

// MakeContrailControllerNodeRoleSlice() makes a slice of ContrailControllerNodeRole
func MakeContrailControllerNodeRoleSlice() []*ContrailControllerNodeRole {
	return []*ContrailControllerNodeRole{}
}
