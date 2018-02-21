
package models
// GlobalQosConfig



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propGlobalQosConfig_display_name int = iota
    propGlobalQosConfig_perms2 int = iota
    propGlobalQosConfig_uuid int = iota
    propGlobalQosConfig_fq_name int = iota
    propGlobalQosConfig_parent_uuid int = iota
    propGlobalQosConfig_parent_type int = iota
    propGlobalQosConfig_id_perms int = iota
    propGlobalQosConfig_annotations int = iota
    propGlobalQosConfig_control_traffic_dscp int = iota
)

// GlobalQosConfig 
type GlobalQosConfig struct {

    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ControlTrafficDSCP *ControlTrafficDscpType `json:"control_traffic_dscp,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`


    ForwardingClasss []*ForwardingClass `json:"forwarding_classs,omitempty"`
    QosConfigs []*QosConfig `json:"qos_configs,omitempty"`
    QosQueues []*QosQueue `json:"qos_queues,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *GlobalQosConfig) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeGlobalQosConfig makes GlobalQosConfig
func MakeGlobalQosConfig() *GlobalQosConfig{
    return &GlobalQosConfig{
    //TODO(nati): Apply default
    ParentType: "",
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        ControlTrafficDSCP: MakeControlTrafficDscpType(),
        ParentUUID: "",
        Perms2: MakePermType2(),
        UUID: "",
        FQName: []string{},
        DisplayName: "",
        
        modified: big.NewInt(0),
    }
}



// MakeGlobalQosConfigSlice makes a slice of GlobalQosConfig
func MakeGlobalQosConfigSlice() []*GlobalQosConfig {
    return []*GlobalQosConfig{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *GlobalQosConfig) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc42024a640])
    fqn := []string{}
    
    fqn = GlobalSystemConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *GlobalQosConfig) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_system_config", "_", "-", -1)
}

func (model *GlobalQosConfig) GetDefaultName() string {
    return strings.Replace("default-global_qos_config", "_", "-", -1)
}

func (model *GlobalQosConfig) GetType() string {
    return strings.Replace("global_qos_config", "_", "-", -1)
}

func (model *GlobalQosConfig) GetFQName() []string {
    return model.FQName
}

func (model *GlobalQosConfig) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *GlobalQosConfig) GetParentType() string {
    return model.ParentType
}

func (model *GlobalQosConfig) GetUuid() string {
    return model.UUID
}

func (model *GlobalQosConfig) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *GlobalQosConfig) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *GlobalQosConfig) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *GlobalQosConfig) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *GlobalQosConfig) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propGlobalQosConfig_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_control_traffic_dscp) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ControlTrafficDSCP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ControlTrafficDSCP as control_traffic_dscp")
        }
        msg["control_traffic_dscp"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propGlobalQosConfig_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *GlobalQosConfig) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *GlobalQosConfig) UpdateReferences() error {
    return nil
}


