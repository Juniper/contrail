package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeOpenstackNetworkNode makes OpenstackNetworkNode
// nolint
func MakeOpenstackNetworkNode() *OpenstackNetworkNode {
	return &OpenstackNetworkNode{
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

// MakeOpenstackNetworkNode makes OpenstackNetworkNode
// nolint
func InterfaceToOpenstackNetworkNode(i interface{}) *OpenstackNetworkNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenstackNetworkNode{
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
	}
}

// MakeOpenstackNetworkNodeSlice() makes a slice of OpenstackNetworkNode
// nolint
func MakeOpenstackNetworkNodeSlice() []*OpenstackNetworkNode {
	return []*OpenstackNetworkNode{}
}

// InterfaceToOpenstackNetworkNodeSlice() makes a slice of OpenstackNetworkNode
// nolint
func InterfaceToOpenstackNetworkNodeSlice(i interface{}) []*OpenstackNetworkNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenstackNetworkNode{}
	for _, item := range list {
		result = append(result, InterfaceToOpenstackNetworkNode(item))
	}
	return result
}
