
package models
// QosQueue



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propQosQueue_min_bandwidth int = iota
    propQosQueue_fq_name int = iota
    propQosQueue_display_name int = iota
    propQosQueue_uuid int = iota
    propQosQueue_annotations int = iota
    propQosQueue_perms2 int = iota
    propQosQueue_qos_queue_identifier int = iota
    propQosQueue_max_bandwidth int = iota
    propQosQueue_parent_uuid int = iota
    propQosQueue_parent_type int = iota
    propQosQueue_id_perms int = iota
)

// QosQueue 
type QosQueue struct {

    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    QosQueueIdentifier int `json:"qos_queue_identifier,omitempty"`
    MaxBandwidth int `json:"max_bandwidth,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    UUID string `json:"uuid,omitempty"`
    MinBandwidth int `json:"min_bandwidth,omitempty"`
    FQName []string `json:"fq_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *QosQueue) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeQosQueue makes QosQueue
func MakeQosQueue() *QosQueue{
    return &QosQueue{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        QosQueueIdentifier: 0,
        MaxBandwidth: 0,
        ParentUUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        MinBandwidth: 0,
        FQName: []string{},
        DisplayName: "",
        UUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeQosQueueSlice makes a slice of QosQueue
func MakeQosQueueSlice() []*QosQueue {
    return []*QosQueue{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *QosQueue) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_qos_config:0xc4202e4aa0])
    fqn := []string{}
    
    fqn = GlobalQosConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *QosQueue) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_qos_config", "_", "-", -1)
}

func (model *QosQueue) GetDefaultName() string {
    return strings.Replace("default-qos_queue", "_", "-", -1)
}

func (model *QosQueue) GetType() string {
    return strings.Replace("qos_queue", "_", "-", -1)
}

func (model *QosQueue) GetFQName() []string {
    return model.FQName
}

func (model *QosQueue) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *QosQueue) GetParentType() string {
    return model.ParentType
}

func (model *QosQueue) GetUuid() string {
    return model.UUID
}

func (model *QosQueue) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *QosQueue) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *QosQueue) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *QosQueue) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *QosQueue) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propQosQueue_min_bandwidth) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MinBandwidth); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MinBandwidth as min_bandwidth")
        }
        msg["min_bandwidth"] = &val
    }
    
    if model.modified.Bit(propQosQueue_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propQosQueue_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propQosQueue_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propQosQueue_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propQosQueue_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propQosQueue_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propQosQueue_qos_queue_identifier) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.QosQueueIdentifier); err != nil {
            return nil, errors.Wrap(err, "Marshal of: QosQueueIdentifier as qos_queue_identifier")
        }
        msg["qos_queue_identifier"] = &val
    }
    
    if model.modified.Bit(propQosQueue_max_bandwidth) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MaxBandwidth); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MaxBandwidth as max_bandwidth")
        }
        msg["max_bandwidth"] = &val
    }
    
    if model.modified.Bit(propQosQueue_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propQosQueue_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *QosQueue) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *QosQueue) UpdateReferences() error {
    return nil
}


