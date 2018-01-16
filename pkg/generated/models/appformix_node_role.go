package models

// AppformixNodeRole

import "encoding/json"

// AppformixNodeRole
type AppformixNodeRole struct {
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
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
		IDPerms:                   MakeIdPermsType(),
		Perms2:                    MakePermType2(),
		ParentUUID:                "",
		ParentType:                "",
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		FQName:      []string{},
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ProvisioningStartTime: "",
	}
}

// MakeAppformixNodeRoleSlice() makes a slice of AppformixNodeRole
func MakeAppformixNodeRoleSlice() []*AppformixNodeRole {
	return []*AppformixNodeRole{}
}
