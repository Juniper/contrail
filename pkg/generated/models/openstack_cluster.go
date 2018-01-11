package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	ProvisioningProgress                      int            `json:"provisioning_progress"`
	OpenstackWebui                            string         `json:"openstack_webui"`
	PublicGateway                             string         `json:"public_gateway"`
	PublicIP                                  string         `json:"public_ip"`
	Perms2                                    *PermType2     `json:"perms2"`
	UUID                                      string         `json:"uuid"`
	DefaultOsdDrives                          string         `json:"default_osd_drives"`
	FQName                                    []string       `json:"fq_name"`
	IDPerms                                   *IdPermsType   `json:"id_perms"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time"`
	DisplayName                               string         `json:"display_name"`
	DefaultJournalDrives                      string         `json:"default_journal_drives"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start"`
	AdminPassword                             string         `json:"admin_password"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end"`
	Annotations                               *KeyValuePairs `json:"annotations"`
	ParentUUID                                string         `json:"parent_uuid"`
	ProvisioningLog                           string         `json:"provisioning_log"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members"`
	ExternalNetCidr                           string         `json:"external_net_cidr"`
	ParentType                                string         `json:"parent_type"`
	ProvisioningState                         string         `json:"provisioning_state"`
	ContrailClusterID                         string         `json:"contrail_cluster_id"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members"`
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
		ParentUUID:                                "",
		AdminPassword:                             "",
		DefaultCapacityDrives:                     "",
		ExternalAllocationPoolEnd:                 "",
		Annotations:                               MakeKeyValuePairs(),
		ProvisioningLog:                           "",
		ProvisioningProgressStage:                 "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalNetCidr:                           "",
		ParentType:                                "",
		ProvisioningState:                         "",
		ContrailClusterID:                         "",
		DefaultStorageAccessBondInterfaceMembers:  "",
		UUID:                        "",
		ProvisioningProgress:        0,
		OpenstackWebui:              "",
		PublicGateway:               "",
		PublicIP:                    "",
		Perms2:                      MakePermType2(),
		DefaultOsdDrives:            "",
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		ProvisioningStartTime:       "",
		ExternalAllocationPoolStart: "",
		DisplayName:                 "",
		DefaultJournalDrives:        "",
		DefaultPerformanceDrives:    "",
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
