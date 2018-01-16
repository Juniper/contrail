package models

// VPNGroup

import "encoding/json"

// VPNGroup
type VPNGroup struct {
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	Type                      string         `json:"type,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`

	LocationRefs []*VPNGroupLocationRef `json:"location_refs,omitempty"`
}

// VPNGroupLocationRef references each other
type VPNGroupLocationRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *VPNGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVPNGroup makes VPNGroup
func MakeVPNGroup() *VPNGroup {
	return &VPNGroup{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		FQName:          []string{},
		DisplayName:     "",
		Perms2:          MakePermType2(),
		ProvisioningLog: "",
		Type:            "",
		ProvisioningStartTime: "",
		ParentType:            "",
		IDPerms:               MakeIdPermsType(),
		Annotations:           MakeKeyValuePairs(),
		ProvisioningState:     "",
	}
}

// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
	return []*VPNGroup{}
}
