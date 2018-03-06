package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeContrailVrouterNode makes ContrailVrouterNode
// nolint
func MakeContrailVrouterNode() *ContrailVrouterNode {
	return &ContrailVrouterNode{
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

// MakeContrailVrouterNode makes ContrailVrouterNode
// nolint
func InterfaceToContrailVrouterNode(i interface{}) *ContrailVrouterNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailVrouterNode{
		//TODO(nati): Apply default
		ProvisioningLog:             common.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:        common.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:   common.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:       common.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:           common.InterfaceToString(m["provisioning_state"]),
		UUID:                        common.InterfaceToString(m["uuid"]),
		ParentUUID:                  common.InterfaceToString(m["parent_uuid"]),
		ParentType:                  common.InterfaceToString(m["parent_type"]),
		FQName:                      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                 common.InterfaceToString(m["display_name"]),
		Annotations:                 InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                      InterfaceToPermType2(m["perms2"]),
		DefaultGateway:              common.InterfaceToString(m["default_gateway"]),
		VrouterBondInterface:        common.InterfaceToString(m["vrouter_bond_interface"]),
		VrouterBondInterfaceMembers: common.InterfaceToString(m["vrouter_bond_interface_members"]),
		VrouterType:                 common.InterfaceToString(m["vrouter_type"]),
	}
}

// MakeContrailVrouterNodeSlice() makes a slice of ContrailVrouterNode
// nolint
func MakeContrailVrouterNodeSlice() []*ContrailVrouterNode {
	return []*ContrailVrouterNode{}
}

// InterfaceToContrailVrouterNodeSlice() makes a slice of ContrailVrouterNode
// nolint
func InterfaceToContrailVrouterNodeSlice(i interface{}) []*ContrailVrouterNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailVrouterNode{}
	for _, item := range list {
		result = append(result, InterfaceToContrailVrouterNode(item))
	}
	return result
}
