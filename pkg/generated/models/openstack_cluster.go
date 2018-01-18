package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DefaultPerformanceDrives                  string         `json:"default_performance_drives,omitempty"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members,omitempty"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end,omitempty"`
	UUID                                      string         `json:"uuid,omitempty"`
	DefaultJournalDrives                      string         `json:"default_journal_drives,omitempty"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start,omitempty"`
	ParentType                                string         `json:"parent_type,omitempty"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives,omitempty"`
	ExternalNetCidr                           string         `json:"external_net_cidr,omitempty"`
	AdminPassword                             string         `json:"admin_password,omitempty"`
	ProvisioningProgress                      int            `json:"provisioning_progress,omitempty"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members,omitempty"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time,omitempty"`
	PublicGateway                             string         `json:"public_gateway,omitempty"`
	PublicIP                                  string         `json:"public_ip,omitempty"`
	ProvisioningLog                           string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState                         string         `json:"provisioning_state,omitempty"`
	OpenstackWebui                            string         `json:"openstack_webui,omitempty"`
	ContrailClusterID                         string         `json:"contrail_cluster_id,omitempty"`
	ParentUUID                                string         `json:"parent_uuid,omitempty"`
	IDPerms                                   *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                                    *PermType2     `json:"perms2,omitempty"`
	FQName                                    []string       `json:"fq_name,omitempty"`
	DisplayName                               string         `json:"display_name,omitempty"`
	Annotations                               *KeyValuePairs `json:"annotations,omitempty"`
	DefaultOsdDrives                          string         `json:"default_osd_drives,omitempty"`
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
		ParentType:                               "",
		DefaultCapacityDrives:                    "",
		ExternalAllocationPoolStart:              "",
		ExternalNetCidr:                          "",
		AdminPassword:                            "",
		DefaultStorageAccessBondInterfaceMembers: "",
		ProvisioningProgress:                     0,
		ProvisioningStartTime:                    "",
		PublicIP:                                 "",
		ProvisioningLog:                          "",
		ProvisioningProgressStage:                "",
		ProvisioningState:                        "",
		OpenstackWebui:                           "",
		PublicGateway:                            "",
		ParentUUID:                               "",
		IDPerms:                                  MakeIdPermsType(),
		Perms2:                                   MakePermType2(),
		ContrailClusterID:                        "",
		DisplayName:                              "",
		Annotations:                              MakeKeyValuePairs(),
		DefaultOsdDrives:                         "",
		FQName:                                   []string{},
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalAllocationPoolEnd:                 "",
		UUID:                     "",
		DefaultJournalDrives:     "",
		DefaultPerformanceDrives: "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
