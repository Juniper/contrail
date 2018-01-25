package models
// OpenstackCluster



import "encoding/json"

// OpenstackCluster 
//proteus:generate
type OpenstackCluster struct {

    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    AdminPassword string `json:"admin_password,omitempty"`
    ContrailClusterID string `json:"contrail_cluster_id,omitempty"`
    DefaultCapacityDrives string `json:"default_capacity_drives,omitempty"`
    DefaultJournalDrives string `json:"default_journal_drives,omitempty"`
    DefaultOsdDrives string `json:"default_osd_drives,omitempty"`
    DefaultPerformanceDrives string `json:"default_performance_drives,omitempty"`
    DefaultStorageAccessBondInterfaceMembers string `json:"default_storage_access_bond_interface_members,omitempty"`
    DefaultStorageBackendBondInterfaceMembers string `json:"default_storage_backend_bond_interface_members,omitempty"`
    ExternalAllocationPoolEnd string `json:"external_allocation_pool_end,omitempty"`
    ExternalAllocationPoolStart string `json:"external_allocation_pool_start,omitempty"`
    ExternalNetCidr string `json:"external_net_cidr,omitempty"`
    OpenstackWebui string `json:"openstack_webui,omitempty"`
    PublicGateway string `json:"public_gateway,omitempty"`
    PublicIP string `json:"public_ip,omitempty"`


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
    ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ProvisioningStartTime: "",
        ProvisioningState: "",
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        AdminPassword: "",
        ContrailClusterID: "",
        DefaultCapacityDrives: "",
        DefaultJournalDrives: "",
        DefaultOsdDrives: "",
        DefaultPerformanceDrives: "",
        DefaultStorageAccessBondInterfaceMembers: "",
        DefaultStorageBackendBondInterfaceMembers: "",
        ExternalAllocationPoolEnd: "",
        ExternalAllocationPoolStart: "",
        ExternalNetCidr: "",
        OpenstackWebui: "",
        PublicGateway: "",
        PublicIP: "",
        
    }
}



// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
    return []*OpenstackCluster{}
}
