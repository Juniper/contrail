package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeKubernetesMasterNode makes KubernetesMasterNode
// nolint
func MakeKubernetesMasterNode() *KubernetesMasterNode {
	return &KubernetesMasterNode{
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

// MakeKubernetesMasterNode makes KubernetesMasterNode
// nolint
func InterfaceToKubernetesMasterNode(i interface{}) *KubernetesMasterNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &KubernetesMasterNode{
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

// MakeKubernetesMasterNodeSlice() makes a slice of KubernetesMasterNode
// nolint
func MakeKubernetesMasterNodeSlice() []*KubernetesMasterNode {
	return []*KubernetesMasterNode{}
}

// InterfaceToKubernetesMasterNodeSlice() makes a slice of KubernetesMasterNode
// nolint
func InterfaceToKubernetesMasterNodeSlice(i interface{}) []*KubernetesMasterNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*KubernetesMasterNode{}
	for _, item := range list {
		result = append(result, InterfaceToKubernetesMasterNode(item))
	}
	return result
}
