package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
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
		PublicGateway:                             "",
		Perms2:                                    MakePermType2(),
		ContrailClusterID:                         "",
		DefaultOsdDrives:                          "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalNetCidr:                           "",
		IDPerms:                                   MakeIdPermsType(),
		UUID:                                      "",
		DefaultPerformanceDrives:                 "",
		ExternalAllocationPoolEnd:                "",
		DisplayName:                              "",
		ProvisioningLog:                          "",
		DefaultJournalDrives:                     "",
		DefaultStorageAccessBondInterfaceMembers: "",
		ExternalAllocationPoolStart:              "",
		ProvisioningStartTime:                    "",
		DefaultCapacityDrives:                    "",
		ParentUUID:                               "",
		ProvisioningProgress:                     0,
		AdminPassword:                            "",
		ParentType:                               "",
		ProvisioningState:                        "",
		FQName:                                   []string{},
		Annotations:                              MakeKeyValuePairs(),
		OpenstackWebui:                           "",
		PublicIP:                                 "",
		ProvisioningProgressStage:                "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
