package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
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
		DefaultPerformanceDrives:                  "",
		DefaultJournalDrives:                      "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalAllocationPoolStart:               "",
		ExternalNetCidr:                           "",
		ProvisioningProgressStage:                 "",
		ProvisioningProgress:                      0,
		DefaultCapacityDrives:                     "",
		ProvisioningState:                         "",
		ProvisioningLog:                           "",
		DefaultStorageAccessBondInterfaceMembers:  "",
		UUID:                      "",
		OpenstackWebui:            "",
		IDPerms:                   MakeIdPermsType(),
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		AdminPassword:             "",
		ExternalAllocationPoolEnd: "",
		PublicGateway:             "",
		PublicIP:                  "",
		FQName:                    []string{},
		DisplayName:               "",
		ProvisioningStartTime:     "",
		ContrailClusterID:         "",
		DefaultOsdDrives:          "",
		ParentUUID:                "",
		ParentType:                "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
