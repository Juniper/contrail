package models

// VPNGroup

import "encoding/json"

// VPNGroup
type VPNGroup struct {
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	Type                      string         `json:"type,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`

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
		Perms2:                    MakePermType2(),
		ProvisioningProgress:      0,
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		Type:                      "",
		ParentUUID:                "",
		ProvisioningState:         "",
		ProvisioningLog:           "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentType:                "",
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
	}
}

// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
	return []*VPNGroup{}
}
