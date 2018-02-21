
package models
// QosConfig



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propQosConfig_fq_name int = iota
    propQosConfig_id_perms int = iota
    propQosConfig_display_name int = iota
    propQosConfig_annotations int = iota
    propQosConfig_perms2 int = iota
    propQosConfig_vlan_priority_entries int = iota
    propQosConfig_default_forwarding_class_id int = iota
    propQosConfig_parent_type int = iota
    propQosConfig_uuid int = iota
    propQosConfig_parent_uuid int = iota
    propQosConfig_qos_config_type int = iota
    propQosConfig_mpls_exp_entries int = iota
    propQosConfig_dscp_entries int = iota
)

// QosConfig 
type QosConfig struct {

    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    QosConfigType QosConfigType `json:"qos_config_type,omitempty"`
    MPLSExpEntries *QosIdForwardingClassPairs `json:"mpls_exp_entries,omitempty"`
    DSCPEntries *QosIdForwardingClassPairs `json:"dscp_entries,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    VlanPriorityEntries *QosIdForwardingClassPairs `json:"vlan_priority_entries,omitempty"`
    DefaultForwardingClassID ForwardingClassId `json:"default_forwarding_class_id,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    GlobalSystemConfigRefs []*QosConfigGlobalSystemConfigRef `json:"global_system_config_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// QosConfigGlobalSystemConfigRef references each other
type QosConfigGlobalSystemConfigRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *QosConfig) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeQosConfig makes QosConfig
func MakeQosConfig() *QosConfig{
    return &QosConfig{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        VlanPriorityEntries: MakeQosIdForwardingClassPairs(),
        DefaultForwardingClassID: MakeForwardingClassId(),
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        QosConfigType: MakeQosConfigType(),
        MPLSExpEntries: MakeQosIdForwardingClassPairs(),
        DSCPEntries: MakeQosIdForwardingClassPairs(),
        UUID: "",
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeQosConfigSlice makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
    return []*QosConfig{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *QosConfig) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_qos_config:0xc4202e4a00 project:0xc4202e4960])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *QosConfig) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *QosConfig) GetDefaultName() string {
    return strings.Replace("default-qos_config", "_", "-", -1)
}

func (model *QosConfig) GetType() string {
    return strings.Replace("qos_config", "_", "-", -1)
}

func (model *QosConfig) GetFQName() []string {
    return model.FQName
}

func (model *QosConfig) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *QosConfig) GetParentType() string {
    return model.ParentType
}

func (model *QosConfig) GetUuid() string {
    return model.UUID
}

func (model *QosConfig) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *QosConfig) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *QosConfig) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *QosConfig) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *QosConfig) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propQosConfig_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propQosConfig_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propQosConfig_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propQosConfig_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propQosConfig_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propQosConfig_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propQosConfig_vlan_priority_entries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VlanPriorityEntries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VlanPriorityEntries as vlan_priority_entries")
        }
        msg["vlan_priority_entries"] = &val
    }
    
    if model.modified.Bit(propQosConfig_default_forwarding_class_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultForwardingClassID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultForwardingClassID as default_forwarding_class_id")
        }
        msg["default_forwarding_class_id"] = &val
    }
    
    if model.modified.Bit(propQosConfig_dscp_entries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSCPEntries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSCPEntries as dscp_entries")
        }
        msg["dscp_entries"] = &val
    }
    
    if model.modified.Bit(propQosConfig_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propQosConfig_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propQosConfig_qos_config_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.QosConfigType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: QosConfigType as qos_config_type")
        }
        msg["qos_config_type"] = &val
    }
    
    if model.modified.Bit(propQosConfig_mpls_exp_entries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MPLSExpEntries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MPLSExpEntries as mpls_exp_entries")
        }
        msg["mpls_exp_entries"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *QosConfig) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *QosConfig) UpdateReferences() error {
    return nil
}


