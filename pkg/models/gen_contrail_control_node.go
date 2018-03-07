package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeContrailControlNode makes ContrailControlNode
// nolint
func MakeContrailControlNode() *ContrailControlNode {
	return &ContrailControlNode{
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

// MakeContrailControlNode makes ContrailControlNode
// nolint
func InterfaceToContrailControlNode(i interface{}) *ContrailControlNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailControlNode{
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

		NodeRefs: InterfaceToContrailControlNodeNodeRefs(m["node_refs"]),
	}
}

func InterfaceToContrailControlNodeNodeRefs(i interface{}) []*ContrailControlNodeNodeRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ContrailControlNodeNodeRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ContrailControlNodeNodeRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeContrailControlNodeSlice() makes a slice of ContrailControlNode
// nolint
func MakeContrailControlNodeSlice() []*ContrailControlNode {
	return []*ContrailControlNode{}
}

// InterfaceToContrailControlNodeSlice() makes a slice of ContrailControlNode
// nolint
func InterfaceToContrailControlNodeSlice(i interface{}) []*ContrailControlNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailControlNode{}
	for _, item := range list {
		result = append(result, InterfaceToContrailControlNode(item))
	}
	return result
}
