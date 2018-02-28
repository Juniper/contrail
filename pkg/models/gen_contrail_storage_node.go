package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeContrailStorageNode makes ContrailStorageNode
// nolint
func MakeContrailStorageNode() *ContrailStorageNode {
	return &ContrailStorageNode{
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
		JournalDrives:                      "",
		OsdDrives:                          "",
		StorageAccessBondInterfaceMembers:  "",
		StorageBackendBondInterfaceMembers: "",
	}
}

// MakeContrailStorageNode makes ContrailStorageNode
// nolint
func InterfaceToContrailStorageNode(i interface{}) *ContrailStorageNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailStorageNode{
		//TODO(nati): Apply default
		ProvisioningLog:                    common.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:               common.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:          common.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:              common.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:                  common.InterfaceToString(m["provisioning_state"]),
		UUID:                               common.InterfaceToString(m["uuid"]),
		ParentUUID:                         common.InterfaceToString(m["parent_uuid"]),
		ParentType:                         common.InterfaceToString(m["parent_type"]),
		FQName:                             common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                        common.InterfaceToString(m["display_name"]),
		Annotations:                        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                             InterfaceToPermType2(m["perms2"]),
		JournalDrives:                      common.InterfaceToString(m["journal_drives"]),
		OsdDrives:                          common.InterfaceToString(m["osd_drives"]),
		StorageAccessBondInterfaceMembers:  common.InterfaceToString(m["storage_access_bond_interface_members"]),
		StorageBackendBondInterfaceMembers: common.InterfaceToString(m["storage_backend_bond_interface_members"]),
	}
}

// MakeContrailStorageNodeSlice() makes a slice of ContrailStorageNode
// nolint
func MakeContrailStorageNodeSlice() []*ContrailStorageNode {
	return []*ContrailStorageNode{}
}

// InterfaceToContrailStorageNodeSlice() makes a slice of ContrailStorageNode
// nolint
func InterfaceToContrailStorageNodeSlice(i interface{}) []*ContrailStorageNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailStorageNode{}
	for _, item := range list {
		result = append(result, InterfaceToContrailStorageNode(item))
	}
	return result
}
