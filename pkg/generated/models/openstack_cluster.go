package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
}

// String returns json representation of the object
func (model *OpenstackCluster) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeOpenstackCluster makes OpenstackCluster
func MakeOpenstackCluster() *OpenstackCluster {
	return &OpenstackCluster{
		//TODO(nati): Apply default
		ProvisioningProgressStage:                 "",
		ExternalAllocationPoolEnd:                 "",
		ProvisioningState:                         "",
		ProvisioningProgress:                      0,
		DefaultOsdDrives:                          "",
		DefaultPerformanceDrives:                  "",
		OpenstackWebui:                            "",
		DisplayName:                               "",
		Perms2:                                    MakePermType2(),
		AdminPassword:                             "",
		DefaultCapacityDrives:                     "",
		ProvisioningStartTime:                     "",
		ContrailClusterID:                         "",
		Annotations:                               MakeKeyValuePairs(),
		ParentUUID:                                "",
		ProvisioningLog:                           "",
		IDPerms:                                   MakeIdPermsType(),
		PublicIP:                                  "",
		DefaultJournalDrives:                      "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalAllocationPoolStart:               "",
		PublicGateway:                             "",
		UUID:                                      "",
		DefaultStorageAccessBondInterfaceMembers: "",
		ExternalNetCidr:                          "",
		ParentType:                               "",
		FQName:                                   []string{},
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
