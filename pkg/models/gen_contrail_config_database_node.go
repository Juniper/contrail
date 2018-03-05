package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeContrailConfigDatabaseNode makes ContrailConfigDatabaseNode
// nolint
func MakeContrailConfigDatabaseNode() *ContrailConfigDatabaseNode {
	return &ContrailConfigDatabaseNode{
		//TODO(nati): Apply default
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		ConfigurationVersion:      0,
	}
}

// MakeContrailConfigDatabaseNode makes ContrailConfigDatabaseNode
// nolint
func InterfaceToContrailConfigDatabaseNode(i interface{}) *ContrailConfigDatabaseNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailConfigDatabaseNode{
		//TODO(nati): Apply default
		ProvisioningLog:           common.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:      common.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage: common.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:     common.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:         common.InterfaceToString(m["provisioning_state"]),
		UUID:                      common.InterfaceToString(m["uuid"]),
		ParentUUID:                common.InterfaceToString(m["parent_uuid"]),
		ParentType:                common.InterfaceToString(m["parent_type"]),
		FQName:                    common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               common.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:      common.InterfaceToInt64(m["configuration_version"]),
	}
}

// MakeContrailConfigDatabaseNodeSlice() makes a slice of ContrailConfigDatabaseNode
// nolint
func MakeContrailConfigDatabaseNodeSlice() []*ContrailConfigDatabaseNode {
	return []*ContrailConfigDatabaseNode{}
}

// InterfaceToContrailConfigDatabaseNodeSlice() makes a slice of ContrailConfigDatabaseNode
// nolint
func InterfaceToContrailConfigDatabaseNodeSlice(i interface{}) []*ContrailConfigDatabaseNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailConfigDatabaseNode{}
	for _, item := range list {
		result = append(result, InterfaceToContrailConfigDatabaseNode(item))
	}
	return result
}
