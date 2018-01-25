package models
// VPNGroup



import "encoding/json"

// VPNGroup 
//proteus:generate
type VPNGroup struct {

    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    Type string `json:"type,omitempty"`

    LocationRefs []*VPNGroupLocationRef `json:"location_refs,omitempty"`

}


// VPNGroupLocationRef references each other
type VPNGroupLocationRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *VPNGroup) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVPNGroup makes VPNGroup
func MakeVPNGroup() *VPNGroup{
    return &VPNGroup{
    //TODO(nati): Apply default
    ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ProvisioningStartTime: "",
        ProvisioningState: "",
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Type: "",
        
    }
}



// MakeVPNGroupSlice() makes a slice of VPNGroup
func MakeVPNGroupSlice() []*VPNGroup {
    return []*VPNGroup{}
}
