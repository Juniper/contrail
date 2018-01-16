package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
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
		DefaultOsdDrives:            "",
		DefaultPerformanceDrives:    "",
		OpenstackWebui:              "",
		ProvisioningProgressStage:   "",
		ProvisioningState:           "",
		DefaultCapacityDrives:       "",
		FQName:                      []string{},
		ProvisioningProgress:        0,
		UUID:                        "",
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		PublicGateway:               "",
		ExternalAllocationPoolEnd:   "",
		ExternalNetCidr:             "",
		ExternalAllocationPoolStart: "",
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		ParentUUID:                  "",
		DefaultStorageAccessBondInterfaceMembers:  "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ContrailClusterID:                         "",
		PublicIP:                                  "",
		ParentType:                                "",
		DefaultJournalDrives:                      "",
		ProvisioningLog:                           "",
		ProvisioningStartTime:                     "",
		AdminPassword:                             "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
