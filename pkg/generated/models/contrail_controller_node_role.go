
package models
// ContrailControllerNodeRole



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propContrailControllerNodeRole_uuid int = iota
    propContrailControllerNodeRole_display_name int = iota
    propContrailControllerNodeRole_provisioning_state int = iota
    propContrailControllerNodeRole_provisioning_log int = iota
    propContrailControllerNodeRole_provisioning_progress int = iota
    propContrailControllerNodeRole_provisioning_progress_stage int = iota
    propContrailControllerNodeRole_perms2 int = iota
    propContrailControllerNodeRole_parent_type int = iota
    propContrailControllerNodeRole_fq_name int = iota
    propContrailControllerNodeRole_id_perms int = iota
    propContrailControllerNodeRole_annotations int = iota
    propContrailControllerNodeRole_provisioning_start_time int = iota
    propContrailControllerNodeRole_parent_uuid int = iota
)

// ContrailControllerNodeRole 
type ContrailControllerNodeRole struct {

    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ContrailControllerNodeRole) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeContrailControllerNodeRole makes ContrailControllerNodeRole
func MakeContrailControllerNodeRole() *ContrailControllerNodeRole{
    return &ContrailControllerNodeRole{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        UUID: "",
        DisplayName: "",
        ProvisioningState: "",
        ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        ProvisioningStartTime: "",
        
        modified: big.NewInt(0),
    }
}



// MakeContrailControllerNodeRoleSlice makes a slice of ContrailControllerNodeRole
func MakeContrailControllerNodeRoleSlice() []*ContrailControllerNodeRole {
    return []*ContrailControllerNodeRole{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ContrailControllerNodeRole) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ContrailControllerNodeRole) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ContrailControllerNodeRole) GetDefaultName() string {
    return strings.Replace("default-contrail_controller_node_role", "_", "-", -1)
}

func (model *ContrailControllerNodeRole) GetType() string {
    return strings.Replace("contrail_controller_node_role", "_", "-", -1)
}

func (model *ContrailControllerNodeRole) GetFQName() []string {
    return model.FQName
}

func (model *ContrailControllerNodeRole) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ContrailControllerNodeRole) GetParentType() string {
    return model.ParentType
}

func (model *ContrailControllerNodeRole) GetUuid() string {
    return model.UUID
}

func (model *ContrailControllerNodeRole) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ContrailControllerNodeRole) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ContrailControllerNodeRole) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ContrailControllerNodeRole) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ContrailControllerNodeRole) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propContrailControllerNodeRole_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propContrailControllerNodeRole_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ContrailControllerNodeRole) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ContrailControllerNodeRole) UpdateReferences() error {
    return nil
}


