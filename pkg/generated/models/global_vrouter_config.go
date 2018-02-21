
package models
// GlobalVrouterConfig



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propGlobalVrouterConfig_enable_security_logging int = iota
    propGlobalVrouterConfig_id_perms int = iota
    propGlobalVrouterConfig_perms2 int = iota
    propGlobalVrouterConfig_uuid int = iota
    propGlobalVrouterConfig_flow_export_rate int = iota
    propGlobalVrouterConfig_fq_name int = iota
    propGlobalVrouterConfig_forwarding_mode int = iota
    propGlobalVrouterConfig_linklocal_services int = iota
    propGlobalVrouterConfig_encapsulation_priorities int = iota
    propGlobalVrouterConfig_parent_uuid int = iota
    propGlobalVrouterConfig_annotations int = iota
    propGlobalVrouterConfig_ecmp_hashing_include_fields int = iota
    propGlobalVrouterConfig_vxlan_network_identifier_mode int = iota
    propGlobalVrouterConfig_parent_type int = iota
    propGlobalVrouterConfig_display_name int = iota
    propGlobalVrouterConfig_flow_aging_timeout_list int = iota
)

// GlobalVrouterConfig 
type GlobalVrouterConfig struct {

    FlowAgingTimeoutList *FlowAgingTimeoutList `json:"flow_aging_timeout_list,omitempty"`
    VxlanNetworkIdentifierMode VxlanNetworkIdentifierModeType `json:"vxlan_network_identifier_mode,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    FlowExportRate int `json:"flow_export_rate,omitempty"`
    EnableSecurityLogging bool `json:"enable_security_logging"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ForwardingMode ForwardingModeType `json:"forwarding_mode,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    EcmpHashingIncludeFields *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
    LinklocalServices *LinklocalServicesTypes `json:"linklocal_services,omitempty"`
    EncapsulationPriorities *EncapsulationPrioritiesType `json:"encapsulation_priorities,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`


    SecurityLoggingObjects []*SecurityLoggingObject `json:"security_logging_objects,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *GlobalVrouterConfig) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeGlobalVrouterConfig makes GlobalVrouterConfig
func MakeGlobalVrouterConfig() *GlobalVrouterConfig{
    return &GlobalVrouterConfig{
    //TODO(nati): Apply default
    ForwardingMode: MakeForwardingModeType(),
        FQName: []string{},
        Annotations: MakeKeyValuePairs(),
        EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
        LinklocalServices: MakeLinklocalServicesTypes(),
        EncapsulationPriorities: MakeEncapsulationPrioritiesType(),
        ParentUUID: "",
        FlowAgingTimeoutList: MakeFlowAgingTimeoutList(),
        VxlanNetworkIdentifierMode: MakeVxlanNetworkIdentifierModeType(),
        ParentType: "",
        DisplayName: "",
        UUID: "",
        FlowExportRate: 0,
        EnableSecurityLogging: false,
        IDPerms: MakeIdPermsType(),
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeGlobalVrouterConfigSlice makes a slice of GlobalVrouterConfig
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
    return []*GlobalVrouterConfig{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *GlobalVrouterConfig) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc42024a820])
    fqn := []string{}
    
    fqn = GlobalSystemConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *GlobalVrouterConfig) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_system_config", "_", "-", -1)
}

func (model *GlobalVrouterConfig) GetDefaultName() string {
    return strings.Replace("default-global_vrouter_config", "_", "-", -1)
}

func (model *GlobalVrouterConfig) GetType() string {
    return strings.Replace("global_vrouter_config", "_", "-", -1)
}

func (model *GlobalVrouterConfig) GetFQName() []string {
    return model.FQName
}

func (model *GlobalVrouterConfig) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *GlobalVrouterConfig) GetParentType() string {
    return model.ParentType
}

func (model *GlobalVrouterConfig) GetUuid() string {
    return model.UUID
}

func (model *GlobalVrouterConfig) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *GlobalVrouterConfig) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *GlobalVrouterConfig) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *GlobalVrouterConfig) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *GlobalVrouterConfig) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propGlobalVrouterConfig_flow_export_rate) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FlowExportRate); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FlowExportRate as flow_export_rate")
        }
        msg["flow_export_rate"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_enable_security_logging) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EnableSecurityLogging); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EnableSecurityLogging as enable_security_logging")
        }
        msg["enable_security_logging"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_forwarding_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ForwardingMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ForwardingMode as forwarding_mode")
        }
        msg["forwarding_mode"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_ecmp_hashing_include_fields) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EcmpHashingIncludeFields); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EcmpHashingIncludeFields as ecmp_hashing_include_fields")
        }
        msg["ecmp_hashing_include_fields"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_linklocal_services) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LinklocalServices); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LinklocalServices as linklocal_services")
        }
        msg["linklocal_services"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_encapsulation_priorities) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EncapsulationPriorities); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EncapsulationPriorities as encapsulation_priorities")
        }
        msg["encapsulation_priorities"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_flow_aging_timeout_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FlowAgingTimeoutList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FlowAgingTimeoutList as flow_aging_timeout_list")
        }
        msg["flow_aging_timeout_list"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_vxlan_network_identifier_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VxlanNetworkIdentifierMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VxlanNetworkIdentifierMode as vxlan_network_identifier_mode")
        }
        msg["vxlan_network_identifier_mode"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propGlobalVrouterConfig_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *GlobalVrouterConfig) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *GlobalVrouterConfig) UpdateReferences() error {
    return nil
}


