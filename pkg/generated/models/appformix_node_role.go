package models

// AppformixNodeRole

import "encoding/json"

// AppformixNodeRole
type AppformixNodeRole struct {
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
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
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ProvisioningStartTime:     "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningState:         "",
		ParentUUID:                "",
		ParentType:                "",
		ProvisioningLog:           "",
	}
}

// MakeAppformixNodeRoleSlice() makes a slice of AppformixNodeRole
func MakeAppformixNodeRoleSlice() []*AppformixNodeRole {
	return []*AppformixNodeRole{}
}
