package models

// VPNGroup

import "encoding/json"

// VPNGroup
type VPNGroup struct {
	ParentUUID                string         `json:"parent_uuid"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage"`
	DisplayName               string         `json:"display_name"`
	Perms2                    *PermType2     `json:"perms2"`
	ParentType                string         `json:"parent_type"`
	ProvisioningLog           string         `json:"provisioning_log"`
	Annotations               *KeyValuePairs `json:"annotations"`
	UUID                      string         `json:"uuid"`
	ProvisioningStartTime     string         `json:"provisioning_start_time"`
	IDPerms                   *IdPermsType   `json:"id_perms"`
	FQName                    []string       `json:"fq_name"`
	ProvisioningProgress      int            `json:"provisioning_progress"`
	Type                      string         `json:"type"`
	ProvisioningState         string         `json:"provisioning_state"`

	LocationRefs []*VPNGroupLocationRef `json:"location_refs"`
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
		ParentUUID:                "",
		ProvisioningProgressStage: "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentType:                "",
		ProvisioningLog:           "",
		IDPerms:                   MakeIdPermsType(),
		FQName:                    []string{},
		ProvisioningStartTime:     "",
		Type:                 "",
		ProvisioningState:    "",
		ProvisioningProgress: 0,
	}
}

// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
	return []*VPNGroup{}
}
