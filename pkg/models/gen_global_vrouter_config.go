package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeGlobalVrouterConfig makes GlobalVrouterConfig
// nolint
func MakeGlobalVrouterConfig() *GlobalVrouterConfig {
	return &GlobalVrouterConfig{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		EcmpHashingIncludeFields:   MakeEcmpHashingIncludeFields(),
		FlowAgingTimeoutList:       MakeFlowAgingTimeoutList(),
		ForwardingMode:             "",
		FlowExportRate:             0,
		LinklocalServices:          MakeLinklocalServicesTypes(),
		EncapsulationPriorities:    MakeEncapsulationPrioritiesType(),
		VxlanNetworkIdentifierMode: "",
		EnableSecurityLogging:      false,
	}
}

// MakeGlobalVrouterConfig makes GlobalVrouterConfig
// nolint
func InterfaceToGlobalVrouterConfig(i interface{}) *GlobalVrouterConfig {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &GlobalVrouterConfig{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		EcmpHashingIncludeFields:   InterfaceToEcmpHashingIncludeFields(m["ecmp_hashing_include_fields"]),
		FlowAgingTimeoutList:       InterfaceToFlowAgingTimeoutList(m["flow_aging_timeout_list"]),
		ForwardingMode:             common.InterfaceToString(m["forwarding_mode"]),
		FlowExportRate:             common.InterfaceToInt64(m["flow_export_rate"]),
		LinklocalServices:          InterfaceToLinklocalServicesTypes(m["linklocal_services"]),
		EncapsulationPriorities:    InterfaceToEncapsulationPrioritiesType(m["encapsulation_priorities"]),
		VxlanNetworkIdentifierMode: common.InterfaceToString(m["vxlan_network_identifier_mode"]),
		EnableSecurityLogging:      common.InterfaceToBool(m["enable_security_logging"]),
	}
}

// MakeGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
// nolint
func MakeGlobalVrouterConfigSlice() []*GlobalVrouterConfig {
	return []*GlobalVrouterConfig{}
}

// InterfaceToGlobalVrouterConfigSlice() makes a slice of GlobalVrouterConfig
// nolint
func InterfaceToGlobalVrouterConfigSlice(i interface{}) []*GlobalVrouterConfig {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*GlobalVrouterConfig{}
	for _, item := range list {
		result = append(result, InterfaceToGlobalVrouterConfig(item))
	}
	return result
}
