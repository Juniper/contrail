package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
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
		ProvisioningState:    "",
		DefaultJournalDrives: "",
		DefaultOsdDrives:     "",
		PublicIP:             "",
		UUID:                 "",
		ProvisioningStartTime:                    "",
		ExternalNetCidr:                          "",
		PublicGateway:                            "",
		DisplayName:                              "",
		ProvisioningProgressStage:                "",
		DefaultStorageAccessBondInterfaceMembers: "",
		ParentUUID: "",
		ParentType: "",
		Perms2:     MakePermType2(),
		ExternalAllocationPoolEnd: "",
		FQName:                                    []string{},
		ProvisioningProgress:                      0,
		ContrailClusterID:                         "",
		DefaultCapacityDrives:                     "",
		DefaultPerformanceDrives:                  "",
		DefaultStorageBackendBondInterfaceMembers: "",
		OpenstackWebui:                            "",
		IDPerms:                                   MakeIdPermsType(),
		ProvisioningLog:                           "",
		AdminPassword:                             "",
		Annotations:                               MakeKeyValuePairs(),
		ExternalAllocationPoolStart:               "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
