package models

// ContrailControllerNodeRole

import "encoding/json"

// ContrailControllerNodeRole
type ContrailControllerNodeRole struct {
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
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
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ProvisioningState:         "",
		ProvisioningLog:           "",
		ProvisioningProgressStage: "",
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		ProvisioningProgress:      0,
		ProvisioningStartTime:     "",
		IDPerms:                   MakeIdPermsType(),
	}
}

// MakeContrailControllerNodeRoleSlice() makes a slice of ContrailControllerNodeRole
func MakeContrailControllerNodeRoleSlice() []*ContrailControllerNodeRole {
	return []*ContrailControllerNodeRole{}
}
