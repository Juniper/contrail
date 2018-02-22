package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeOpenstackCluster makes OpenstackCluster
func MakeOpenstackCluster() *OpenstackCluster {
	return &OpenstackCluster{
		//TODO(nati): Apply default
		ProvisioningLog:                           "",
		ProvisioningProgress:                      0,
		ProvisioningProgressStage:                 "",
		ProvisioningStartTime:                     "",
		ProvisioningState:                         "",
		UUID:                                      "",
		ParentUUID:                                "",
		ParentType:                                "",
		FQName:                                    []string{},
		IDPerms:                                   MakeIdPermsType(),
		DisplayName:                               "",
		Annotations:                               MakeKeyValuePairs(),
		Perms2:                                    MakePermType2(),
		AdminPassword:                             "",
		ContrailClusterID:                         "",
		DefaultCapacityDrives:                     "",
		DefaultJournalDrives:                      "",
		DefaultOsdDrives:                          "",
		DefaultPerformanceDrives:                  "",
		DefaultStorageAccessBondInterfaceMembers:  "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ExternalAllocationPoolEnd:                 "",
		ExternalAllocationPoolStart:               "",
		ExternalNetCidr:                           "",
		OpenstackWebui:                            "",
		PublicGateway:                             "",
		PublicIP:                                  "",
	}
}

// MakeOpenstackCluster makes OpenstackCluster
func InterfaceToOpenstackCluster(i interface{}) *OpenstackCluster {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenstackCluster{
		//TODO(nati): Apply default
		ProvisioningLog:                           schema.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:                      schema.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:                 schema.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:                     schema.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:                         schema.InterfaceToString(m["provisioning_state"]),
		UUID:                                      schema.InterfaceToString(m["uuid"]),
		ParentUUID:                                schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                                schema.InterfaceToString(m["parent_type"]),
		FQName:                                    schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                               schema.InterfaceToString(m["display_name"]),
		Annotations:                               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                                    InterfaceToPermType2(m["perms2"]),
		AdminPassword:                             schema.InterfaceToString(m["admin_password"]),
		ContrailClusterID:                         schema.InterfaceToString(m["contrail_cluster_id"]),
		DefaultCapacityDrives:                     schema.InterfaceToString(m["default_capacity_drives"]),
		DefaultJournalDrives:                      schema.InterfaceToString(m["default_journal_drives"]),
		DefaultOsdDrives:                          schema.InterfaceToString(m["default_osd_drives"]),
		DefaultPerformanceDrives:                  schema.InterfaceToString(m["default_performance_drives"]),
		DefaultStorageAccessBondInterfaceMembers:  schema.InterfaceToString(m["default_storage_access_bond_interface_members"]),
		DefaultStorageBackendBondInterfaceMembers: schema.InterfaceToString(m["default_storage_backend_bond_interface_members"]),
		ExternalAllocationPoolEnd:                 schema.InterfaceToString(m["external_allocation_pool_end"]),
		ExternalAllocationPoolStart:               schema.InterfaceToString(m["external_allocation_pool_start"]),
		ExternalNetCidr:                           schema.InterfaceToString(m["external_net_cidr"]),
		OpenstackWebui:                            schema.InterfaceToString(m["openstack_webui"]),
		PublicGateway:                             schema.InterfaceToString(m["public_gateway"]),
		PublicIP:                                  schema.InterfaceToString(m["public_ip"]),
	}
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}

// InterfaceToOpenstackClusterSlice() makes a slice of OpenstackCluster
func InterfaceToOpenstackClusterSlice(i interface{}) []*OpenstackCluster {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenstackCluster{}
	for _, item := range list {
		result = append(result, InterfaceToOpenstackCluster(item))
	}
	return result
}
