
package models
// OpenstackComputeNodeRole



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propOpenstackComputeNodeRole_vrouter_bond_interface int = iota
    propOpenstackComputeNodeRole_display_name int = iota
    propOpenstackComputeNodeRole_perms2 int = iota
    propOpenstackComputeNodeRole_uuid int = iota
    propOpenstackComputeNodeRole_provisioning_progress int = iota
    propOpenstackComputeNodeRole_provisioning_progress_stage int = iota
    propOpenstackComputeNodeRole_vrouter_bond_interface_members int = iota
    propOpenstackComputeNodeRole_vrouter_type int = iota
    propOpenstackComputeNodeRole_parent_type int = iota
    propOpenstackComputeNodeRole_fq_name int = iota
    propOpenstackComputeNodeRole_provisioning_state int = iota
    propOpenstackComputeNodeRole_parent_uuid int = iota
    propOpenstackComputeNodeRole_id_perms int = iota
    propOpenstackComputeNodeRole_annotations int = iota
    propOpenstackComputeNodeRole_provisioning_log int = iota
    propOpenstackComputeNodeRole_default_gateway int = iota
    propOpenstackComputeNodeRole_provisioning_start_time int = iota
)

// OpenstackComputeNodeRole 
type OpenstackComputeNodeRole struct {

    UUID string `json:"uuid,omitempty"`
    VrouterBondInterface string `json:"vrouter_bond_interface,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    VrouterBondInterfaceMembers string `json:"vrouter_bond_interface_members,omitempty"`
    VrouterType string `json:"vrouter_type,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    DefaultGateway string `json:"default_gateway,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *OpenstackComputeNodeRole) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeOpenstackComputeNodeRole makes OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRole() *OpenstackComputeNodeRole{
    return &OpenstackComputeNodeRole{
    //TODO(nati): Apply default
    ProvisioningProgressStage: "",
        VrouterBondInterfaceMembers: "",
        VrouterType: "",
        ParentType: "",
        FQName: []string{},
        ProvisioningState: "",
        ProvisioningProgress: 0,
        ParentUUID: "",
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        ProvisioningLog: "",
        DefaultGateway: "",
        ProvisioningStartTime: "",
        VrouterBondInterface: "",
        DisplayName: "",
        Perms2: MakePermType2(),
        UUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeOpenstackComputeNodeRoleSlice makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
    return []*OpenstackComputeNodeRole{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *OpenstackComputeNodeRole) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *OpenstackComputeNodeRole) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *OpenstackComputeNodeRole) GetDefaultName() string {
    return strings.Replace("default-openstack_compute_node_role", "_", "-", -1)
}

func (model *OpenstackComputeNodeRole) GetType() string {
    return strings.Replace("openstack_compute_node_role", "_", "-", -1)
}

func (model *OpenstackComputeNodeRole) GetFQName() []string {
    return model.FQName
}

func (model *OpenstackComputeNodeRole) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *OpenstackComputeNodeRole) GetParentType() string {
    return model.ParentType
}

func (model *OpenstackComputeNodeRole) GetUuid() string {
    return model.UUID
}

func (model *OpenstackComputeNodeRole) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *OpenstackComputeNodeRole) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *OpenstackComputeNodeRole) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *OpenstackComputeNodeRole) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *OpenstackComputeNodeRole) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propOpenstackComputeNodeRole_vrouter_bond_interface) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VrouterBondInterface); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VrouterBondInterface as vrouter_bond_interface")
        }
        msg["vrouter_bond_interface"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_vrouter_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VrouterBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VrouterBondInterfaceMembers as vrouter_bond_interface_members")
        }
        msg["vrouter_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_vrouter_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VrouterType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VrouterType as vrouter_type")
        }
        msg["vrouter_type"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_default_gateway) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultGateway); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultGateway as default_gateway")
        }
        msg["default_gateway"] = &val
    }
    
    if model.modified.Bit(propOpenstackComputeNodeRole_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *OpenstackComputeNodeRole) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *OpenstackComputeNodeRole) UpdateReferences() error {
    return nil
}


