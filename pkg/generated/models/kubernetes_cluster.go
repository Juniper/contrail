
package models
// KubernetesCluster



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propKubernetesCluster_perms2 int = iota
    propKubernetesCluster_uuid int = iota
    propKubernetesCluster_parent_uuid int = iota
    propKubernetesCluster_parent_type int = iota
    propKubernetesCluster_contrail_cluster_id int = iota
    propKubernetesCluster_kuberunetes_dashboard int = iota
    propKubernetesCluster_fq_name int = iota
    propKubernetesCluster_annotations int = iota
    propKubernetesCluster_id_perms int = iota
    propKubernetesCluster_display_name int = iota
)

// KubernetesCluster 
type KubernetesCluster struct {

    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ContrailClusterID string `json:"contrail_cluster_id,omitempty"`
    KuberunetesDashboard string `json:"kuberunetes_dashboard,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *KubernetesCluster) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeKubernetesCluster makes KubernetesCluster
func MakeKubernetesCluster() *KubernetesCluster{
    return &KubernetesCluster{
    //TODO(nati): Apply default
    IDPerms: MakeIdPermsType(),
        DisplayName: "",
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        ContrailClusterID: "",
        KuberunetesDashboard: "",
        FQName: []string{},
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeKubernetesClusterSlice makes a slice of KubernetesCluster
func MakeKubernetesClusterSlice() []*KubernetesCluster {
    return []*KubernetesCluster{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *KubernetesCluster) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *KubernetesCluster) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *KubernetesCluster) GetDefaultName() string {
    return strings.Replace("default-kubernetes_cluster", "_", "-", -1)
}

func (model *KubernetesCluster) GetType() string {
    return strings.Replace("kubernetes_cluster", "_", "-", -1)
}

func (model *KubernetesCluster) GetFQName() []string {
    return model.FQName
}

func (model *KubernetesCluster) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *KubernetesCluster) GetParentType() string {
    return model.ParentType
}

func (model *KubernetesCluster) GetUuid() string {
    return model.UUID
}

func (model *KubernetesCluster) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *KubernetesCluster) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *KubernetesCluster) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *KubernetesCluster) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *KubernetesCluster) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propKubernetesCluster_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_contrail_cluster_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ContrailClusterID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ContrailClusterID as contrail_cluster_id")
        }
        msg["contrail_cluster_id"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_kuberunetes_dashboard) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.KuberunetesDashboard); err != nil {
            return nil, errors.Wrap(err, "Marshal of: KuberunetesDashboard as kuberunetes_dashboard")
        }
        msg["kuberunetes_dashboard"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propKubernetesCluster_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *KubernetesCluster) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *KubernetesCluster) UpdateReferences() error {
    return nil
}


