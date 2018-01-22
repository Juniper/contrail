package models

// AppformixNodeRole

import "encoding/json"

// AppformixNodeRole
type AppformixNodeRole struct {
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
}

// String returns json representation of the object
func (model *AppformixNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAppformixNodeRole makes AppformixNodeRole
func MakeAppformixNodeRole() *AppformixNodeRole {
	return &AppformixNodeRole{
		//TODO(nati): Apply default
		ProvisioningState:         "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		ProvisioningLog:           "",
		ProvisioningStartTime:     "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
	}
}

// MakeAppformixNodeRoleSlice() makes a slice of AppformixNodeRole
func MakeAppformixNodeRoleSlice() []*AppformixNodeRole {
	return []*AppformixNodeRole{}
}
