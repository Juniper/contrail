package models
// ContrailAnalyticsDatabaseNodeRole



import "encoding/json"

// ContrailAnalyticsDatabaseNodeRole 
//proteus:generate
type ContrailAnalyticsDatabaseNodeRole struct {

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


}



// String returns json representation of the object
func (model *ContrailAnalyticsDatabaseNodeRole) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeContrailAnalyticsDatabaseNodeRole makes ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRole() *ContrailAnalyticsDatabaseNodeRole{
    return &ContrailAnalyticsDatabaseNodeRole{
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
        
    }
}



// MakeContrailAnalyticsDatabaseNodeRoleSlice() makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
    return []*ContrailAnalyticsDatabaseNodeRole{}
}
