
package models
// KubernetesNode



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propKubernetesNode_annotations int = iota
    propKubernetesNode_perms2 int = iota
    propKubernetesNode_uuid int = iota
    propKubernetesNode_parent_uuid int = iota
    propKubernetesNode_parent_type int = iota
    propKubernetesNode_provisioning_log int = iota
    propKubernetesNode_provisioning_start_time int = iota
    propKubernetesNode_display_name int = iota
    propKubernetesNode_fq_name int = iota
    propKubernetesNode_provisioning_progress int = iota
    propKubernetesNode_provisioning_progress_stage int = iota
    propKubernetesNode_provisioning_state int = iota
    propKubernetesNode_id_perms int = iota
)

// KubernetesNode 
type KubernetesNode struct {

    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *KubernetesNode) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeKubernetesNode makes KubernetesNode
func MakeKubernetesNode() *KubernetesNode{
    return &KubernetesNode{
    //TODO(nati): Apply default
    ProvisioningLog: "",
        ProvisioningStartTime: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        FQName: []string{},
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ProvisioningState: "",
        
        modified: big.NewInt(0),
    }
}



// MakeKubernetesNodeSlice makes a slice of KubernetesNode
func MakeKubernetesNodeSlice() []*KubernetesNode {
    return []*KubernetesNode{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *KubernetesNode) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *KubernetesNode) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *KubernetesNode) GetDefaultName() string {
    return strings.Replace("default-kubernetes_node", "_", "-", -1)
}

func (model *KubernetesNode) GetType() string {
    return strings.Replace("kubernetes_node", "_", "-", -1)
}

func (model *KubernetesNode) GetFQName() []string {
    return model.FQName
}

func (model *KubernetesNode) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *KubernetesNode) GetParentType() string {
    return model.ParentType
}

func (model *KubernetesNode) GetUuid() string {
    return model.UUID
}

func (model *KubernetesNode) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *KubernetesNode) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *KubernetesNode) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *KubernetesNode) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *KubernetesNode) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propKubernetesNode_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propKubernetesNode_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *KubernetesNode) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *KubernetesNode) UpdateReferences() error {
    return nil
}


