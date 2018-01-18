package models

// VPNGroup

import "encoding/json"

// VPNGroup
type VPNGroup struct {
	DisplayName               string         `json:"display_name,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	Type                      string         `json:"type,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	UUID                      string         `json:"uuid,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`

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
		ParentType: "",
		IDPerms:    MakeIdPermsType(),
		Perms2:     MakePermType2(),
		ProvisioningProgressStage: "",
		UUID:                  "",
		FQName:                []string{},
		DisplayName:           "",
		ProvisioningProgress:  0,
		ProvisioningLog:       "",
		Type:                  "",
		ParentUUID:            "",
		Annotations:           MakeKeyValuePairs(),
		ProvisioningStartTime: "",
		ProvisioningState:     "",
	}
}

// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
	return []*VPNGroup{}
}
