
package models
// ContrailCluster



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propContrailCluster_data_ttl int = iota
    propContrailCluster_default_gateway int = iota
    propContrailCluster_flow_ttl int = iota
    propContrailCluster_display_name int = iota
    propContrailCluster_default_vrouter_bond_interface_members int = iota
    propContrailCluster_annotations int = iota
    propContrailCluster_fq_name int = iota
    propContrailCluster_config_audit_ttl int = iota
    propContrailCluster_contrail_webui int = iota
    propContrailCluster_default_vrouter_bond_interface int = iota
    propContrailCluster_statistics_ttl int = iota
    propContrailCluster_perms2 int = iota
    propContrailCluster_uuid int = iota
    propContrailCluster_parent_uuid int = iota
    propContrailCluster_parent_type int = iota
    propContrailCluster_id_perms int = iota
)

// ContrailCluster 
type ContrailCluster struct {

    DefaultVrouterBondInterfaceMembers string `json:"default_vrouter_bond_interface_members,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ConfigAuditTTL string `json:"config_audit_ttl,omitempty"`
    ContrailWebui string `json:"contrail_webui,omitempty"`
    DefaultVrouterBondInterface string `json:"default_vrouter_bond_interface,omitempty"`
    StatisticsTTL string `json:"statistics_ttl,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    DataTTL string `json:"data_ttl,omitempty"`
    DefaultGateway string `json:"default_gateway,omitempty"`
    FlowTTL string `json:"flow_ttl,omitempty"`
    DisplayName string `json:"display_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ContrailCluster) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeContrailCluster makes ContrailCluster
func MakeContrailCluster() *ContrailCluster{
    return &ContrailCluster{
    //TODO(nati): Apply default
    ParentUUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        Perms2: MakePermType2(),
        UUID: "",
        FlowTTL: "",
        DisplayName: "",
        DataTTL: "",
        DefaultGateway: "",
        FQName: []string{},
        DefaultVrouterBondInterfaceMembers: "",
        Annotations: MakeKeyValuePairs(),
        DefaultVrouterBondInterface: "",
        StatisticsTTL: "",
        ConfigAuditTTL: "",
        ContrailWebui: "",
        
        modified: big.NewInt(0),
    }
}



// MakeContrailClusterSlice makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
    return []*ContrailCluster{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ContrailCluster) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ContrailCluster) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ContrailCluster) GetDefaultName() string {
    return strings.Replace("default-contrail_cluster", "_", "-", -1)
}

func (model *ContrailCluster) GetType() string {
    return strings.Replace("contrail_cluster", "_", "-", -1)
}

func (model *ContrailCluster) GetFQName() []string {
    return model.FQName
}

func (model *ContrailCluster) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ContrailCluster) GetParentType() string {
    return model.ParentType
}

func (model *ContrailCluster) GetUuid() string {
    return model.UUID
}

func (model *ContrailCluster) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ContrailCluster) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ContrailCluster) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ContrailCluster) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ContrailCluster) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propContrailCluster_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_data_ttl) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DataTTL); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DataTTL as data_ttl")
        }
        msg["data_ttl"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_default_gateway) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultGateway); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultGateway as default_gateway")
        }
        msg["default_gateway"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_flow_ttl) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FlowTTL); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FlowTTL as flow_ttl")
        }
        msg["flow_ttl"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_default_vrouter_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultVrouterBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultVrouterBondInterfaceMembers as default_vrouter_bond_interface_members")
        }
        msg["default_vrouter_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_config_audit_ttl) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ConfigAuditTTL); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ConfigAuditTTL as config_audit_ttl")
        }
        msg["config_audit_ttl"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_contrail_webui) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ContrailWebui); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ContrailWebui as contrail_webui")
        }
        msg["contrail_webui"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_default_vrouter_bond_interface) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultVrouterBondInterface); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultVrouterBondInterface as default_vrouter_bond_interface")
        }
        msg["default_vrouter_bond_interface"] = &val
    }
    
    if model.modified.Bit(propContrailCluster_statistics_ttl) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StatisticsTTL); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StatisticsTTL as statistics_ttl")
        }
        msg["statistics_ttl"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ContrailCluster) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ContrailCluster) UpdateReferences() error {
    return nil
}


