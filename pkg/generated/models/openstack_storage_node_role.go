
package models
// OpenstackStorageNodeRole



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propOpenstackStorageNodeRole_storage_backend_bond_interface_members int = iota
    propOpenstackStorageNodeRole_id_perms int = iota
    propOpenstackStorageNodeRole_storage_access_bond_interface_members int = iota
    propOpenstackStorageNodeRole_parent_type int = iota
    propOpenstackStorageNodeRole_fq_name int = iota
    propOpenstackStorageNodeRole_provisioning_progress_stage int = iota
    propOpenstackStorageNodeRole_provisioning_start_time int = iota
    propOpenstackStorageNodeRole_journal_drives int = iota
    propOpenstackStorageNodeRole_annotations int = iota
    propOpenstackStorageNodeRole_perms2 int = iota
    propOpenstackStorageNodeRole_provisioning_log int = iota
    propOpenstackStorageNodeRole_provisioning_state int = iota
    propOpenstackStorageNodeRole_osd_drives int = iota
    propOpenstackStorageNodeRole_parent_uuid int = iota
    propOpenstackStorageNodeRole_display_name int = iota
    propOpenstackStorageNodeRole_uuid int = iota
    propOpenstackStorageNodeRole_provisioning_progress int = iota
)

// OpenstackStorageNodeRole 
type OpenstackStorageNodeRole struct {

    UUID string `json:"uuid,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    OsdDrives string `json:"osd_drives,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    StorageBackendBondInterfaceMembers string `json:"storage_backend_bond_interface_members,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    StorageAccessBondInterfaceMembers string `json:"storage_access_bond_interface_members,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    JournalDrives string `json:"journal_drives,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *OpenstackStorageNodeRole) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeOpenstackStorageNodeRole makes OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRole() *OpenstackStorageNodeRole{
    return &OpenstackStorageNodeRole{
    //TODO(nati): Apply default
    ProvisioningProgress: 0,
        OsdDrives: "",
        ParentUUID: "",
        DisplayName: "",
        UUID: "",
        StorageBackendBondInterfaceMembers: "",
        IDPerms: MakeIdPermsType(),
        ProvisioningStartTime: "",
        StorageAccessBondInterfaceMembers: "",
        ParentType: "",
        FQName: []string{},
        ProvisioningProgressStage: "",
        ProvisioningState: "",
        JournalDrives: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ProvisioningLog: "",
        
        modified: big.NewInt(0),
    }
}



// MakeOpenstackStorageNodeRoleSlice makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
    return []*OpenstackStorageNodeRole{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *OpenstackStorageNodeRole) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *OpenstackStorageNodeRole) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *OpenstackStorageNodeRole) GetDefaultName() string {
    return strings.Replace("default-openstack_storage_node_role", "_", "-", -1)
}

func (model *OpenstackStorageNodeRole) GetType() string {
    return strings.Replace("openstack_storage_node_role", "_", "-", -1)
}

func (model *OpenstackStorageNodeRole) GetFQName() []string {
    return model.FQName
}

func (model *OpenstackStorageNodeRole) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *OpenstackStorageNodeRole) GetParentType() string {
    return model.ParentType
}

func (model *OpenstackStorageNodeRole) GetUuid() string {
    return model.UUID
}

func (model *OpenstackStorageNodeRole) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *OpenstackStorageNodeRole) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *OpenstackStorageNodeRole) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *OpenstackStorageNodeRole) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *OpenstackStorageNodeRole) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propOpenstackStorageNodeRole_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_storage_backend_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StorageBackendBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StorageBackendBondInterfaceMembers as storage_backend_bond_interface_members")
        }
        msg["storage_backend_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_storage_access_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StorageAccessBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StorageAccessBondInterfaceMembers as storage_access_bond_interface_members")
        }
        msg["storage_access_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_journal_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.JournalDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: JournalDrives as journal_drives")
        }
        msg["journal_drives"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propOpenstackStorageNodeRole_osd_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OsdDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OsdDrives as osd_drives")
        }
        msg["osd_drives"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *OpenstackStorageNodeRole) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *OpenstackStorageNodeRole) UpdateReferences() error {
    return nil
}


