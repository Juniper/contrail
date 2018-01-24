package models

// AppformixNodeRole

import "encoding/json"

// AppformixNodeRole
type AppformixNodeRole struct {
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
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
		Perms2:                    MakePermType2(),
		UUID:                      "",
		FQName:                    []string{},
		Annotations:               MakeKeyValuePairs(),
		ProvisioningProgress:      0,
		ProvisioningStartTime:     "",
		ParentUUID:                "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ProvisioningLog:           "",
		ProvisioningProgressStage: "",
	}
}

// MakeAppformixNodeRoleSlice() makes a slice of AppformixNodeRole
func MakeAppformixNodeRoleSlice() []*AppformixNodeRole {
	return []*AppformixNodeRole{}
}
