package models

// AppformixNodeRole

import "encoding/json"

// AppformixNodeRole
type AppformixNodeRole struct {
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
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
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		Perms2:                    MakePermType2(),
		ParentType:                "",
		FQName:                    []string{},
		ProvisioningLog:           "",
		DisplayName:               "",
		ProvisioningProgress:      0,
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentUUID:                "",
		IDPerms:                   MakeIdPermsType(),
	}
}

// MakeAppformixNodeRoleSlice() makes a slice of AppformixNodeRole
func MakeAppformixNodeRoleSlice() []*AppformixNodeRole {
	return []*AppformixNodeRole{}
}
