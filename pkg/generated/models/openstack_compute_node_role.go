package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeOpenstackComputeNodeRole makes OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRole() *OpenstackComputeNodeRole {
	return &OpenstackComputeNodeRole{
		//TODO(nati): Apply default
		ProvisioningLog:             "",
		ProvisioningProgress:        0,
		ProvisioningProgressStage:   "",
		ProvisioningStartTime:       "",
		ProvisioningState:           "",
		UUID:                        "",
		ParentUUID:                  "",
		ParentType:                  "",
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		DefaultGateway:              "",
		VrouterBondInterface:        "",
		VrouterBondInterfaceMembers: "",
		VrouterType:                 "",
	}
}

// MakeOpenstackComputeNodeRole makes OpenstackComputeNodeRole
func InterfaceToOpenstackComputeNodeRole(i interface{}) *OpenstackComputeNodeRole {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &OpenstackComputeNodeRole{
		//TODO(nati): Apply default
		ProvisioningLog:             schema.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:        schema.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:   schema.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:       schema.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:           schema.InterfaceToString(m["provisioning_state"]),
		UUID:                        schema.InterfaceToString(m["uuid"]),
		ParentUUID:                  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                  schema.InterfaceToString(m["parent_type"]),
		FQName:                      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                 schema.InterfaceToString(m["display_name"]),
		Annotations:                 InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                      InterfaceToPermType2(m["perms2"]),
		DefaultGateway:              schema.InterfaceToString(m["default_gateway"]),
		VrouterBondInterface:        schema.InterfaceToString(m["vrouter_bond_interface"]),
		VrouterBondInterfaceMembers: schema.InterfaceToString(m["vrouter_bond_interface_members"]),
		VrouterType:                 schema.InterfaceToString(m["vrouter_type"]),
	}
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}

// InterfaceToOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func InterfaceToOpenstackComputeNodeRoleSlice(i interface{}) []*OpenstackComputeNodeRole {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*OpenstackComputeNodeRole{}
	for _, item := range list {
		result = append(result, InterfaceToOpenstackComputeNodeRole(item))
	}
	return result
}
