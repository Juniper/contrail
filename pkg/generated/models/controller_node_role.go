package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeControllerNodeRole makes ControllerNodeRole
func MakeControllerNodeRole() *ControllerNodeRole {
	return &ControllerNodeRole{
		//TODO(nati): Apply default
		ProvisioningLog:                       "",
		ProvisioningProgress:                  0,
		ProvisioningProgressStage:             "",
		ProvisioningStartTime:                 "",
		ProvisioningState:                     "",
		UUID:                                  "",
		ParentUUID:                            "",
		ParentType:                            "",
		FQName:                                []string{},
		IDPerms:                               MakeIdPermsType(),
		DisplayName:                           "",
		Annotations:                           MakeKeyValuePairs(),
		Perms2:                                MakePermType2(),
		CapacityDrives:                        "",
		InternalapiBondInterfaceMembers:       "",
		PerformanceDrives:                     "",
		StorageManagementBondInterfaceMembers: "",
	}
}

// MakeControllerNodeRole makes ControllerNodeRole
func InterfaceToControllerNodeRole(i interface{}) *ControllerNodeRole {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ControllerNodeRole{
		//TODO(nati): Apply default
		ProvisioningLog:                       schema.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:                  schema.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:             schema.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:                 schema.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:                     schema.InterfaceToString(m["provisioning_state"]),
		UUID:                                  schema.InterfaceToString(m["uuid"]),
		ParentUUID:                            schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                            schema.InterfaceToString(m["parent_type"]),
		FQName:                                schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                               InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                           schema.InterfaceToString(m["display_name"]),
		Annotations:                           InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                                InterfaceToPermType2(m["perms2"]),
		CapacityDrives:                        schema.InterfaceToString(m["capacity_drives"]),
		InternalapiBondInterfaceMembers:       schema.InterfaceToString(m["internalapi_bond_interface_members"]),
		PerformanceDrives:                     schema.InterfaceToString(m["performance_drives"]),
		StorageManagementBondInterfaceMembers: schema.InterfaceToString(m["storage_management_bond_interface_members"]),
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}

// InterfaceToControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func InterfaceToControllerNodeRoleSlice(i interface{}) []*ControllerNodeRole {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ControllerNodeRole{}
	for _, item := range list {
		result = append(result, InterfaceToControllerNodeRole(item))
	}
	return result
}
