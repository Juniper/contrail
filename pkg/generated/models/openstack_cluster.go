package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
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
		DefaultOsdDrives:                          "",
		ParentType:                                "",
		ProvisioningStartTime:                     "",
		DefaultStorageAccessBondInterfaceMembers:  "",
		ExternalNetCidr:                           "",
		ProvisioningProgressStage:                 "",
		DefaultPerformanceDrives:                  "",
		PublicIP:                                  "",
		ParentUUID:                                "",
		DisplayName:                               "",
		Annotations:                               MakeKeyValuePairs(),
		ProvisioningLog:                           "",
		ContrailClusterID:                         "",
		UUID:                                      "",
		IDPerms:                                   MakeIdPermsType(),
		ExternalAllocationPoolStart:               "",
		PublicGateway:                             "",
		FQName:                                    []string{},
		Perms2:                                    MakePermType2(),
		ProvisioningProgress:                      0,
		AdminPassword:                             "",
		DefaultCapacityDrives:                     "",
		DefaultJournalDrives:                      "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalAllocationPoolEnd:                 "",
		OpenstackWebui:                            "",
		ProvisioningState:                         "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
