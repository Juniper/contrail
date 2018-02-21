
package models
// OpenstackCluster



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propOpenstackCluster_display_name int = iota
    propOpenstackCluster_provisioning_progress_stage int = iota
    propOpenstackCluster_contrail_cluster_id int = iota
    propOpenstackCluster_default_journal_drives int = iota
    propOpenstackCluster_default_osd_drives int = iota
    propOpenstackCluster_perms2 int = iota
    propOpenstackCluster_public_gateway int = iota
    propOpenstackCluster_default_storage_access_bond_interface_members int = iota
    propOpenstackCluster_external_net_cidr int = iota
    propOpenstackCluster_public_ip int = iota
    propOpenstackCluster_uuid int = iota
    propOpenstackCluster_external_allocation_pool_start int = iota
    propOpenstackCluster_parent_uuid int = iota
    propOpenstackCluster_provisioning_state int = iota
    propOpenstackCluster_parent_type int = iota
    propOpenstackCluster_provisioning_start_time int = iota
    propOpenstackCluster_default_capacity_drives int = iota
    propOpenstackCluster_default_storage_backend_bond_interface_members int = iota
    propOpenstackCluster_id_perms int = iota
    propOpenstackCluster_admin_password int = iota
    propOpenstackCluster_external_allocation_pool_end int = iota
    propOpenstackCluster_openstack_webui int = iota
    propOpenstackCluster_annotations int = iota
    propOpenstackCluster_default_performance_drives int = iota
    propOpenstackCluster_fq_name int = iota
    propOpenstackCluster_provisioning_log int = iota
    propOpenstackCluster_provisioning_progress int = iota
)

// OpenstackCluster 
type OpenstackCluster struct {

    ExternalAllocationPoolStart string `json:"external_allocation_pool_start,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    DefaultCapacityDrives string `json:"default_capacity_drives,omitempty"`
    DefaultStorageBackendBondInterfaceMembers string `json:"default_storage_backend_bond_interface_members,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    AdminPassword string `json:"admin_password,omitempty"`
    ExternalAllocationPoolEnd string `json:"external_allocation_pool_end,omitempty"`
    OpenstackWebui string `json:"openstack_webui,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    DefaultPerformanceDrives string `json:"default_performance_drives,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ContrailClusterID string `json:"contrail_cluster_id,omitempty"`
    DefaultJournalDrives string `json:"default_journal_drives,omitempty"`
    DefaultOsdDrives string `json:"default_osd_drives,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    PublicGateway string `json:"public_gateway,omitempty"`
    DefaultStorageAccessBondInterfaceMembers string `json:"default_storage_access_bond_interface_members,omitempty"`
    ExternalNetCidr string `json:"external_net_cidr,omitempty"`
    PublicIP string `json:"public_ip,omitempty"`
    UUID string `json:"uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *OpenstackCluster) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeOpenstackCluster makes OpenstackCluster
func MakeOpenstackCluster() *OpenstackCluster{
    return &OpenstackCluster{
    //TODO(nati): Apply default
    DefaultPerformanceDrives: "",
        FQName: []string{},
        ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ContrailClusterID: "",
        DefaultJournalDrives: "",
        DefaultOsdDrives: "",
        Perms2: MakePermType2(),
        DisplayName: "",
        PublicGateway: "",
        DefaultStorageAccessBondInterfaceMembers: "",
        ExternalNetCidr: "",
        PublicIP: "",
        UUID: "",
        ExternalAllocationPoolStart: "",
        ParentUUID: "",
        ProvisioningState: "",
        ProvisioningStartTime: "",
        ParentType: "",
        DefaultCapacityDrives: "",
        DefaultStorageBackendBondInterfaceMembers: "",
        IDPerms: MakeIdPermsType(),
        AdminPassword: "",
        ExternalAllocationPoolEnd: "",
        OpenstackWebui: "",
        Annotations: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeOpenstackClusterSlice makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
    return []*OpenstackCluster{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *OpenstackCluster) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *OpenstackCluster) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *OpenstackCluster) GetDefaultName() string {
    return strings.Replace("default-openstack_cluster", "_", "-", -1)
}

func (model *OpenstackCluster) GetType() string {
    return strings.Replace("openstack_cluster", "_", "-", -1)
}

func (model *OpenstackCluster) GetFQName() []string {
    return model.FQName
}

func (model *OpenstackCluster) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *OpenstackCluster) GetParentType() string {
    return model.ParentType
}

func (model *OpenstackCluster) GetUuid() string {
    return model.UUID
}

func (model *OpenstackCluster) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *OpenstackCluster) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *OpenstackCluster) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *OpenstackCluster) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *OpenstackCluster) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propOpenstackCluster_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_default_storage_backend_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultStorageBackendBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultStorageBackendBondInterfaceMembers as default_storage_backend_bond_interface_members")
        }
        msg["default_storage_backend_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_default_capacity_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultCapacityDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultCapacityDrives as default_capacity_drives")
        }
        msg["default_capacity_drives"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_external_allocation_pool_end) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExternalAllocationPoolEnd); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExternalAllocationPoolEnd as external_allocation_pool_end")
        }
        msg["external_allocation_pool_end"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_openstack_webui) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OpenstackWebui); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OpenstackWebui as openstack_webui")
        }
        msg["openstack_webui"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_admin_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AdminPassword); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AdminPassword as admin_password")
        }
        msg["admin_password"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_default_performance_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultPerformanceDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultPerformanceDrives as default_performance_drives")
        }
        msg["default_performance_drives"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_default_journal_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultJournalDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultJournalDrives as default_journal_drives")
        }
        msg["default_journal_drives"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_default_osd_drives) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultOsdDrives); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultOsdDrives as default_osd_drives")
        }
        msg["default_osd_drives"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_contrail_cluster_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ContrailClusterID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ContrailClusterID as contrail_cluster_id")
        }
        msg["contrail_cluster_id"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_public_gateway) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PublicGateway); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PublicGateway as public_gateway")
        }
        msg["public_gateway"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_external_net_cidr) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExternalNetCidr); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExternalNetCidr as external_net_cidr")
        }
        msg["external_net_cidr"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_public_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PublicIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PublicIP as public_ip")
        }
        msg["public_ip"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_default_storage_access_bond_interface_members) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultStorageAccessBondInterfaceMembers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultStorageAccessBondInterfaceMembers as default_storage_access_bond_interface_members")
        }
        msg["default_storage_access_bond_interface_members"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propOpenstackCluster_external_allocation_pool_start) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExternalAllocationPoolStart); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExternalAllocationPoolStart as external_allocation_pool_start")
        }
        msg["external_allocation_pool_start"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *OpenstackCluster) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *OpenstackCluster) UpdateReferences() error {
    return nil
}


