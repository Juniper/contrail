package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeContrailAnalyticsNode makes ContrailAnalyticsNode
func MakeContrailAnalyticsNode() *ContrailAnalyticsNode {
	return &ContrailAnalyticsNode{
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
	}
}

// MakeContrailAnalyticsNode makes ContrailAnalyticsNode
func InterfaceToContrailAnalyticsNode(i interface{}) *ContrailAnalyticsNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailAnalyticsNode{
		//TODO(nati): Apply default
		ProvisioningLog:           schema.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:      schema.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage: schema.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:     schema.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:         schema.InterfaceToString(m["provisioning_state"]),
		UUID:                      schema.InterfaceToString(m["uuid"]),
		ParentUUID:                schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                schema.InterfaceToString(m["parent_type"]),
		FQName:                    schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               schema.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
	}
}

// MakeContrailAnalyticsNodeSlice() makes a slice of ContrailAnalyticsNode
func MakeContrailAnalyticsNodeSlice() []*ContrailAnalyticsNode {
	return []*ContrailAnalyticsNode{}
}

// InterfaceToContrailAnalyticsNodeSlice() makes a slice of ContrailAnalyticsNode
func InterfaceToContrailAnalyticsNodeSlice(i interface{}) []*ContrailAnalyticsNode {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailAnalyticsNode{}
	for _, item := range list {
		result = append(result, InterfaceToContrailAnalyticsNode(item))
	}
	return result
}
