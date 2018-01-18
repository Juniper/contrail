package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
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
		ParentType:                                "",
		ExternalNetCidr:                           "",
		Annotations:                               MakeKeyValuePairs(),
		Perms2:                                    MakePermType2(),
		UUID:                                      "",
		AdminPassword:                             "",
		ProvisioningProgress:                      0,
		ProvisioningStartTime:                     "",
		ContrailClusterID:                         "",
		ParentUUID:                                "",
		ProvisioningProgressStage:                 "",
		PublicGateway:                             "",
		DisplayName:                               "",
		IDPerms:                                   MakeIdPermsType(),
		DefaultPerformanceDrives:                  "",
		DefaultStorageAccessBondInterfaceMembers:  "",
		DefaultStorageBackendBondInterfaceMembers: "",
		OpenstackWebui:                            "",
		DefaultCapacityDrives:                     "",
		DefaultOsdDrives:                          "",
		FQName:                                    []string{},
		ExternalAllocationPoolEnd:   "",
		ProvisioningLog:             "",
		ProvisioningState:           "",
		DefaultJournalDrives:        "",
		ExternalAllocationPoolStart: "",
		PublicIP:                    "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
