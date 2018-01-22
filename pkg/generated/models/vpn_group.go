package models

// VPNGroup

import "encoding/json"

// VPNGroup
type VPNGroup struct {
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	Type                      string         `json:"type,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`

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
		DisplayName:               "",
		ParentUUID:                "",
		ProvisioningProgress:      0,
		Type:                      "",
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ProvisioningLog:           "",
		ParentType:                "",
		FQName:                    []string{},
		Annotations:               MakeKeyValuePairs(),
		IDPerms:                   MakeIdPermsType(),
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
	}
}

// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
	return []*VPNGroup{}
}
