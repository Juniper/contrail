package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeOpenstackStorageNode makes OpenstackStorageNode
func MakeOpenstackStorageNode() *OpenstackStorageNode {
	return &OpenstackStorageNode{
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

// MakeOpenstackStorageNode makes OpenstackStorageNode
func InterfaceToOpenstackStorageNode(i interface{}) *OpenstackStorageNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenstackStorageNode{
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

// MakeOpenstackStorageNodeSlice() makes a slice of OpenstackStorageNode
func MakeOpenstackStorageNodeSlice() []*OpenstackStorageNode {
	return []*OpenstackStorageNode{}
}

// InterfaceToOpenstackStorageNodeSlice() makes a slice of OpenstackStorageNode
func InterfaceToOpenstackStorageNodeSlice(i interface{}) []*OpenstackStorageNode {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenstackStorageNode{}
	for _, item := range list {
		result = append(result, InterfaceToOpenstackStorageNode(item))
	}
	return result
}
