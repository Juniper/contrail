
package models
// AppformixNodeRole



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAppformixNodeRole_uuid int = iota
    propAppformixNodeRole_parent_type int = iota
    propAppformixNodeRole_perms2 int = iota
    propAppformixNodeRole_provisioning_log int = iota
    propAppformixNodeRole_provisioning_progress_stage int = iota
    propAppformixNodeRole_provisioning_start_time int = iota
    propAppformixNodeRole_provisioning_state int = iota
    propAppformixNodeRole_parent_uuid int = iota
    propAppformixNodeRole_fq_name int = iota
    propAppformixNodeRole_id_perms int = iota
    propAppformixNodeRole_display_name int = iota
    propAppformixNodeRole_annotations int = iota
    propAppformixNodeRole_provisioning_progress int = iota
)

// AppformixNodeRole 
type AppformixNodeRole struct {

    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AppformixNodeRole) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAppformixNodeRole makes AppformixNodeRole
func MakeAppformixNodeRole() *AppformixNodeRole{
    return &AppformixNodeRole{
    //TODO(nati): Apply default
    ProvisioningLog: "",
        UUID: "",
        ParentType: "",
        Perms2: MakePermType2(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ProvisioningProgressStage: "",
        ProvisioningStartTime: "",
        ProvisioningState: "",
        ParentUUID: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        ProvisioningProgress: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeAppformixNodeRoleSlice makes a slice of AppformixNodeRole
func MakeAppformixNodeRoleSlice() []*AppformixNodeRole {
    return []*AppformixNodeRole{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AppformixNodeRole) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *AppformixNodeRole) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AppformixNodeRole) GetDefaultName() string {
    return strings.Replace("default-appformix_node_role", "_", "-", -1)
}

func (model *AppformixNodeRole) GetType() string {
    return strings.Replace("appformix_node_role", "_", "-", -1)
}

func (model *AppformixNodeRole) GetFQName() []string {
    return model.FQName
}

func (model *AppformixNodeRole) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AppformixNodeRole) GetParentType() string {
    return model.ParentType
}

func (model *AppformixNodeRole) GetUuid() string {
    return model.UUID
}

func (model *AppformixNodeRole) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AppformixNodeRole) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AppformixNodeRole) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AppformixNodeRole) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AppformixNodeRole) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAppformixNodeRole_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAppformixNodeRole_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AppformixNodeRole) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AppformixNodeRole) UpdateReferences() error {
    return nil
}


