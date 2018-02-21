
package models
// ContrailAnalyticsNode



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propContrailAnalyticsNode_parent_type int = iota
    propContrailAnalyticsNode_provisioning_log int = iota
    propContrailAnalyticsNode_provisioning_progress int = iota
    propContrailAnalyticsNode_provisioning_progress_stage int = iota
    propContrailAnalyticsNode_perms2 int = iota
    propContrailAnalyticsNode_uuid int = iota
    propContrailAnalyticsNode_parent_uuid int = iota
    propContrailAnalyticsNode_fq_name int = iota
    propContrailAnalyticsNode_provisioning_state int = iota
    propContrailAnalyticsNode_provisioning_start_time int = iota
    propContrailAnalyticsNode_id_perms int = iota
    propContrailAnalyticsNode_display_name int = iota
    propContrailAnalyticsNode_annotations int = iota
)

// ContrailAnalyticsNode 
type ContrailAnalyticsNode struct {

    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ContrailAnalyticsNode) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeContrailAnalyticsNode makes ContrailAnalyticsNode
func MakeContrailAnalyticsNode() *ContrailAnalyticsNode{
    return &ContrailAnalyticsNode{
    //TODO(nati): Apply default
    DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        FQName: []string{},
        ProvisioningState: "",
        ProvisioningStartTime: "",
        IDPerms: MakeIdPermsType(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeContrailAnalyticsNodeSlice makes a slice of ContrailAnalyticsNode
func MakeContrailAnalyticsNodeSlice() []*ContrailAnalyticsNode {
    return []*ContrailAnalyticsNode{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ContrailAnalyticsNode) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ContrailAnalyticsNode) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ContrailAnalyticsNode) GetDefaultName() string {
    return strings.Replace("default-contrail_analytics_node", "_", "-", -1)
}

func (model *ContrailAnalyticsNode) GetType() string {
    return strings.Replace("contrail_analytics_node", "_", "-", -1)
}

func (model *ContrailAnalyticsNode) GetFQName() []string {
    return model.FQName
}

func (model *ContrailAnalyticsNode) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ContrailAnalyticsNode) GetParentType() string {
    return model.ParentType
}

func (model *ContrailAnalyticsNode) GetUuid() string {
    return model.UUID
}

func (model *ContrailAnalyticsNode) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ContrailAnalyticsNode) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ContrailAnalyticsNode) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ContrailAnalyticsNode) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ContrailAnalyticsNode) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propContrailAnalyticsNode_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propContrailAnalyticsNode_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ContrailAnalyticsNode) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ContrailAnalyticsNode) UpdateReferences() error {
    return nil
}


