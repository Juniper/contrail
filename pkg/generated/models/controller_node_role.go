
package models
// ControllerNodeRole



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propControllerNodeRole_internalapi_bond_interface_members int = iota
    propControllerNodeRole_storage_management_bond_interface_members int = iota
    propControllerNodeRole_parent_uuid int = iota
    propControllerNodeRole_provisioning_progress int = iota
    propControllerNodeRole_fq_name int = iota
    propControllerNodeRole_id_perms int = iota
    propControllerNodeRole_annotations int = iota
    propControllerNodeRole_provisioning_progress_stage int = iota
    propControllerNodeRole_provisioning_state int = iota
    propControllerNodeRole_capacity_drives int = iota
    propControllerNodeRole_performance_drives int = iota
    propControllerNodeRole_parent_type int = iota
    propControllerNodeRole_provisioning_start_time int = iota
    propControllerNodeRole_provisioning_log int = iota
    propControllerNodeRole_display_name int = iota
    propControllerNodeRole_perms2 int = iota
    propControllerNodeRole_uuid int = iota
)

// ControllerNodeRole 
type ControllerNodeRole struct {

    PerformanceDrives string `json:"performance_drives,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    CapacityDrives string `json:"capacity_drives,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    StorageManagementBondInterfaceMembers string `json:"storage_management_bond_interface_members,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    InternalapiBondInterfaceMembers string `json:"internalapi_bond_interface_members,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    FQName []string `json:"fq_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ControllerNodeRole) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeControllerNodeRole makes ControllerNodeRole
func MakeControllerNodeRole() *ControllerNodeRole{
    return &ControllerNodeRole{
    //TODO(nati): Apply default
    FQName: []string{},
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        ProvisioningProgressStage: "",
        ProvisioningState: "",
        ProvisioningProgress: 0,
        CapacityDrives: "",
        PerformanceDrives: "",
        ParentType: "",
        ProvisioningStartTime: "",
        ProvisioningLog: "",
        DisplayName: "",
        Perms2: MakePermType2(),
        UUID: "",
        InternalapiBondInterfaceMembers: "",
        StorageManagementBondInterfaceMembers: "",
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeControllerNodeRoleSlice makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
    return []*ControllerNodeRole{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ControllerNodeRole) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *ControllerNodeRole) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ControllerNodeRole) GetDefaultName() string {
    return strings.Replace("default-controller_node_role", "_", "-", -1)
}

func (model *ControllerNodeRole) GetType() string {
    return strings.Replace("controller_node_role", "_", "-", -1)
}

func (model *ControllerNodeRole) GetFQName() []string {
    return model.FQName
}

func (model *ControllerNodeRole) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ControllerNodeRole) GetParentType() string {
    return model.ParentType
}

func (model *ControllerNodeRole) GetUuid() string {
    return model.UUID
}

func (model *ControllerNodeRole) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ControllerNodeRole) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ControllerNodeRole) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ControllerNodeRole) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ControllerNodeRole) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propControllerNodeRole_internalapi_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InternalapiBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InternalapiBondInterfaceMembers as internalapi_bond_interface_members")
        }
        msg["internalapi_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_storage_management_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StorageManagementBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StorageManagementBondInterfaceMembers as storage_management_bond_interface_members")
        }
        msg["storage_management_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_capacity_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.CapacityDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: CapacityDrives as capacity_drives")
        }
        msg["capacity_drives"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_performance_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PerformanceDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PerformanceDrives as performance_drives")
        }
        msg["performance_drives"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propControllerNodeRole_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ControllerNodeRole) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ControllerNodeRole) UpdateReferences() error {
    return nil
}


