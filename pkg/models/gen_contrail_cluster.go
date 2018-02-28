package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeContrailCluster makes ContrailCluster
// nolint
func MakeContrailCluster() *ContrailCluster {
	return &ContrailCluster{
		//TODO(nati): Apply default
		ProvisioningLog:                    "",
		ProvisioningProgress:               0,
		ProvisioningProgressStage:          "",
		ProvisioningStartTime:              "",
		ProvisioningState:                  "",
		UUID:                               "",
		ParentUUID:                         "",
		ParentType:                         "",
		FQName:                             []string{},
		IDPerms:                            MakeIdPermsType(),
		DisplayName:                        "",
		Annotations:                        MakeKeyValuePairs(),
		Perms2:                             MakePermType2(),
		ProvisionerType:                    "",
		Orchestrator:                       "",
		Kubernetes:                         "",
		KubernetesMaster:                   "",
		Openstack:                          "",
		ConfigAuditTTL:                     "",
		ContrailConfig:                     "",
		ContrailConfigdb:                   "",
		ContrailControl:                    "",
		ContrailWebui:                      "",
		ContrailAnalytics:                  "",
		ContrailAnalyticsdb:                "",
		ContrailVrouter:                    "",
		DataTTL:                            "",
		DefaultGateway:                     "",
		DefaultVrouterBondInterface:        "",
		DefaultVrouterBondInterfaceMembers: "",
		FlowTTL:       "",
		StatisticsTTL: "",
	}
}

// MakeContrailCluster makes ContrailCluster
// nolint
func InterfaceToContrailCluster(i interface{}) *ContrailCluster {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailCluster{
		//TODO(nati): Apply default
		ProvisioningLog:                    common.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:               common.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:          common.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:              common.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:                  common.InterfaceToString(m["provisioning_state"]),
		UUID:                               common.InterfaceToString(m["uuid"]),
		ParentUUID:                         common.InterfaceToString(m["parent_uuid"]),
		ParentType:                         common.InterfaceToString(m["parent_type"]),
		FQName:                             common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                        common.InterfaceToString(m["display_name"]),
		Annotations:                        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                             InterfaceToPermType2(m["perms2"]),
		ProvisionerType:                    common.InterfaceToString(m["provisioner_type"]),
		Orchestrator:                       common.InterfaceToString(m["orchestrator"]),
		Kubernetes:                         common.InterfaceToString(m["kubernetes"]),
		KubernetesMaster:                   common.InterfaceToString(m["kubernetes_master"]),
		Openstack:                          common.InterfaceToString(m["openstack"]),
		ConfigAuditTTL:                     common.InterfaceToString(m["config_audit_ttl"]),
		ContrailConfig:                     common.InterfaceToString(m["contrail_config"]),
		ContrailConfigdb:                   common.InterfaceToString(m["contrail_configdb"]),
		ContrailControl:                    common.InterfaceToString(m["contrail_control"]),
		ContrailWebui:                      common.InterfaceToString(m["contrail_webui"]),
		ContrailAnalytics:                  common.InterfaceToString(m["contrail_analytics"]),
		ContrailAnalyticsdb:                common.InterfaceToString(m["contrail_analyticsdb"]),
		ContrailVrouter:                    common.InterfaceToString(m["contrail_vrouter"]),
		DataTTL:                            common.InterfaceToString(m["data_ttl"]),
		DefaultGateway:                     common.InterfaceToString(m["default_gateway"]),
		DefaultVrouterBondInterface:        common.InterfaceToString(m["default_vrouter_bond_interface"]),
		DefaultVrouterBondInterfaceMembers: common.InterfaceToString(m["default_vrouter_bond_interface_members"]),
		FlowTTL:       common.InterfaceToString(m["flow_ttl"]),
		StatisticsTTL: common.InterfaceToString(m["statistics_ttl"]),
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
// nolint
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}

// InterfaceToContrailClusterSlice() makes a slice of ContrailCluster
// nolint
func InterfaceToContrailClusterSlice(i interface{}) []*ContrailCluster {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailCluster{}
	for _, item := range list {
		result = append(result, InterfaceToContrailCluster(item))
	}
	return result
}
