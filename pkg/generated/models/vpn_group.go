package models

// VPNGroup

import "encoding/json"

// VPNGroup
type VPNGroup struct {
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	Type                      string         `json:"type,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`

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
		FQName:                    []string{},
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ProvisioningProgress:      0,
		ProvisioningStartTime:     "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		Perms2:                    MakePermType2(),
		ParentUUID:                "",
		ProvisioningProgressStage: "",
		ProvisioningLog:           "",
		Type:                      "",
		DisplayName:               "",
		ProvisioningState:         "",
	}
}

// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
	return []*VPNGroup{}
}
