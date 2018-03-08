package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeOpenstackMonitoringNode makes OpenstackMonitoringNode
// nolint
func MakeOpenstackMonitoringNode() *OpenstackMonitoringNode {
	return &OpenstackMonitoringNode{
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

// MakeOpenstackMonitoringNode makes OpenstackMonitoringNode
// nolint
func InterfaceToOpenstackMonitoringNode(i interface{}) *OpenstackMonitoringNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenstackMonitoringNode{
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

		NodeRefs: InterfaceToOpenstackMonitoringNodeNodeRefs(m["node_refs"]),
	}
}

func InterfaceToOpenstackMonitoringNodeNodeRefs(i interface{}) []*OpenstackMonitoringNodeNodeRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*OpenstackMonitoringNodeNodeRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &OpenstackMonitoringNodeNodeRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeOpenstackMonitoringNodeSlice() makes a slice of OpenstackMonitoringNode
// nolint
func MakeOpenstackMonitoringNodeSlice() []*OpenstackMonitoringNode {
	return []*OpenstackMonitoringNode{}
}

// InterfaceToOpenstackMonitoringNodeSlice() makes a slice of OpenstackMonitoringNode
// nolint
func InterfaceToOpenstackMonitoringNodeSlice(i interface{}) []*OpenstackMonitoringNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenstackMonitoringNode{}
	for _, item := range list {
		result = append(result, InterfaceToOpenstackMonitoringNode(item))
	}
	return result
}
