package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeContrailCluster makes ContrailCluster
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
func InterfaceToContrailCluster(i interface{}) *ContrailCluster {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailCluster{
		//TODO(nati): Apply default
		ProvisioningLog:                    schema.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:               schema.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:          schema.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:              schema.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:                  schema.InterfaceToString(m["provisioning_state"]),
		UUID:                               schema.InterfaceToString(m["uuid"]),
		ParentUUID:                         schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                         schema.InterfaceToString(m["parent_type"]),
		FQName:                             schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                        schema.InterfaceToString(m["display_name"]),
		Annotations:                        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                             InterfaceToPermType2(m["perms2"]),
		ProvisionerType:                    schema.InterfaceToString(m["provisioner_type"]),
		Orchestrator:                       schema.InterfaceToString(m["orchestrator"]),
		Kubernetes:                         schema.InterfaceToString(m["kubernetes"]),
		KubernetesMaster:                   schema.InterfaceToString(m["kubernetes_master"]),
		Openstack:                          schema.InterfaceToString(m["openstack"]),
		ConfigAuditTTL:                     schema.InterfaceToString(m["config_audit_ttl"]),
		ContrailConfig:                     schema.InterfaceToString(m["contrail_config"]),
		ContrailConfigdb:                   schema.InterfaceToString(m["contrail_configdb"]),
		ContrailControl:                    schema.InterfaceToString(m["contrail_control"]),
		ContrailWebui:                      schema.InterfaceToString(m["contrail_webui"]),
		ContrailAnalytics:                  schema.InterfaceToString(m["contrail_analytics"]),
		ContrailAnalyticsdb:                schema.InterfaceToString(m["contrail_analyticsdb"]),
		ContrailVrouter:                    schema.InterfaceToString(m["contrail_vrouter"]),
		DataTTL:                            schema.InterfaceToString(m["data_ttl"]),
		DefaultGateway:                     schema.InterfaceToString(m["default_gateway"]),
		DefaultVrouterBondInterface:        schema.InterfaceToString(m["default_vrouter_bond_interface"]),
		DefaultVrouterBondInterfaceMembers: schema.InterfaceToString(m["default_vrouter_bond_interface_members"]),
		FlowTTL:       schema.InterfaceToString(m["flow_ttl"]),
		StatisticsTTL: schema.InterfaceToString(m["statistics_ttl"]),
	}
}

// MakeContrailClusterSlice() makes a slice of ContrailCluster
func MakeContrailClusterSlice() []*ContrailCluster {
	return []*ContrailCluster{}
}

// InterfaceToContrailClusterSlice() makes a slice of ContrailCluster
func InterfaceToContrailClusterSlice(i interface{}) []*ContrailCluster {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailCluster{}
	for _, item := range list {
		result = append(result, InterfaceToContrailCluster(item))
	}
	return result
}
