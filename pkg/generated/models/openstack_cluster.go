package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DisplayName                               string         `json:"display_name,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
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
		DefaultPerformanceDrives:                 "",
		ContrailClusterID:                        "",
		ProvisioningLog:                          "",
		DisplayName:                              "",
		OpenstackWebui:                           "",
		PublicGateway:                            "",
		UUID:                                     "",
		IDPerms:                                  MakeIdPermsType(),
		ProvisioningProgress:                     0,
		ProvisioningStartTime:                    "",
		ProvisioningState:                        "",
		DefaultStorageAccessBondInterfaceMembers: "",
		DefaultJournalDrives:                     "",
		ExternalAllocationPoolEnd:                "",
		PublicIP:                                 "",
		Perms2:                                   MakePermType2(),
		ExternalAllocationPoolStart: "",
		ParentUUID:                  "",
		ParentType:                  "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalNetCidr:                           "",
		ProvisioningProgressStage:                 "",
		DefaultOsdDrives:                          "",
		Annotations:                               MakeKeyValuePairs(),
		DefaultCapacityDrives:                     "",
		FQName:                                    []string{},
		AdminPassword:                             "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
