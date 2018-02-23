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
		UUID:                               "",
		ParentUUID:                         "",
		ParentType:                         "",
		FQName:                             []string{},
		IDPerms:                            MakeIdPermsType(),
		DisplayName:                        "",
		Annotations:                        MakeKeyValuePairs(),
		Perms2:                             MakePermType2(),
		ConfigAuditTTL:                     "",
		ContrailWebui:                      "",
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
		UUID:                               schema.InterfaceToString(m["uuid"]),
		ParentUUID:                         schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                         schema.InterfaceToString(m["parent_type"]),
		FQName:                             schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                        schema.InterfaceToString(m["display_name"]),
		Annotations:                        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                             InterfaceToPermType2(m["perms2"]),
		ConfigAuditTTL:                     schema.InterfaceToString(m["config_audit_ttl"]),
		ContrailWebui:                      schema.InterfaceToString(m["contrail_webui"]),
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
