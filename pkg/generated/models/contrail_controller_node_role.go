package models

// ContrailControllerNodeRole

import "encoding/json"

// ContrailControllerNodeRole
type ContrailControllerNodeRole struct {
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
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
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		UUID:                      "",
		ProvisioningState:         "",
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ParentUUID:                "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ProvisioningStartTime:     "",
		ProvisioningProgressStage: "",
		FQName: []string{},
	}
}

// MakeContrailControllerNodeRoleSlice() makes a slice of ContrailControllerNodeRole
func MakeContrailControllerNodeRoleSlice() []*ContrailControllerNodeRole {
	return []*ContrailControllerNodeRole{}
}
