package models

// VPNGroup

import "encoding/json"

// VPNGroup
type VPNGroup struct {
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	Type                      string         `json:"type,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`

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
		Perms2:                    MakePermType2(),
		IDPerms:                   MakeIdPermsType(),
		ProvisioningState:         "",
		ProvisioningLog:           "",
		Type:                      "",
		UUID:                      "",
		ParentUUID:                "",
		FQName:                    []string{},
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ParentType:                "",
	}
}

// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
	return []*VPNGroup{}
}
