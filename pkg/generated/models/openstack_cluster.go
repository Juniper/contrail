package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
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
		OpenstackWebui:       "",
		PublicIP:             "",
		ProvisioningProgress: 0,
		ExternalNetCidr:      "",
		FQName:               []string{},
		DefaultPerformanceDrives:    "",
		ExternalAllocationPoolStart: "",
		Annotations:                 MakeKeyValuePairs(),
		ProvisioningLog:             "",
		AdminPassword:               "",
		DefaultOsdDrives:            "",
		ExternalAllocationPoolEnd:   "",
		Perms2:                                   MakePermType2(),
		ProvisioningState:                        "",
		DefaultJournalDrives:                     "",
		ParentType:                               "",
		DisplayName:                              "",
		DefaultStorageAccessBondInterfaceMembers: "",
		UUID:                                      "",
		ParentUUID:                                "",
		IDPerms:                                   MakeIdPermsType(),
		ProvisioningStartTime:                     "",
		DefaultStorageBackendBondInterfaceMembers: "",
		DefaultCapacityDrives:                     "",
		PublicGateway:                             "",
		ProvisioningProgressStage:                 "",
		ContrailClusterID:                         "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
