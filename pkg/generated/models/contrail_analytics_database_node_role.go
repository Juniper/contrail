
package models
// ContrailAnalyticsDatabaseNodeRole



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propContrailAnalyticsDatabaseNodeRole_perms2 int = iota
    propContrailAnalyticsDatabaseNodeRole_fq_name int = iota
    propContrailAnalyticsDatabaseNodeRole_provisioning_progress int = iota
    propContrailAnalyticsDatabaseNodeRole_provisioning_start_time int = iota
    propContrailAnalyticsDatabaseNodeRole_provisioning_state int = iota
    propContrailAnalyticsDatabaseNodeRole_id_perms int = iota
    propContrailAnalyticsDatabaseNodeRole_display_name int = iota
    propContrailAnalyticsDatabaseNodeRole_parent_uuid int = iota
    propContrailAnalyticsDatabaseNodeRole_parent_type int = iota
    propContrailAnalyticsDatabaseNodeRole_provisioning_log int = iota
    propContrailAnalyticsDatabaseNodeRole_provisioning_progress_stage int = iota
    propContrailAnalyticsDatabaseNodeRole_annotations int = iota
    propContrailAnalyticsDatabaseNodeRole_uuid int = iota
)

// ContrailAnalyticsDatabaseNodeRole 
type ContrailAnalyticsDatabaseNodeRole struct {

    Perms2 *PermType2 `json:"perms2,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
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
    ProvisioningState: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Perms2: MakePermType2(),
        FQName: []string{},
        ProvisioningProgress: 0,
        ProvisioningStartTime: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        ProvisioningLog: "",
        ProvisioningProgressStage: "",
        
        modified: big.NewInt(0),
    }
}



// MakeContrailAnalyticsDatabaseNodeRoleSlice makes a slice of ContrailAnalyticsDatabaseNodeRole
func MakeContrailAnalyticsDatabaseNodeRoleSlice() []*ContrailAnalyticsDatabaseNodeRole {
    return []*ContrailAnalyticsDatabaseNodeRole{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ContrailAnalyticsDatabaseNodeRole) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetDefaultName() string {
    return strings.Replace("default-contrail_analytics_database_node_role", "_", "-", -1)
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetType() string {
    return strings.Replace("contrail_analytics_database_node_role", "_", "-", -1)
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetFQName() []string {
    return model.FQName
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetParentType() string {
    return model.ParentType
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetUuid() string {
    return model.UUID
}

func (model *ContrailAnalyticsDatabaseNodeRole) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ContrailAnalyticsDatabaseNodeRole) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ContrailAnalyticsDatabaseNodeRole) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ContrailAnalyticsDatabaseNodeRole) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ContrailAnalyticsDatabaseNodeRole) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsDatabaseNodeRole_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ContrailAnalyticsDatabaseNodeRole) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ContrailAnalyticsDatabaseNodeRole) UpdateReferences() error {
    return nil
}


