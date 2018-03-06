package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBaremetalNode makes BaremetalNode
// nolint
func MakeBaremetalNode() *BaremetalNode {
	return &BaremetalNode{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		Name:                 "",
		DriverInfo:           MakeDriverInfo(),
		BMProperties:         MakeBaremetalProperties(),
		InstanceUUID:         "",
		InstanceInfo:         MakeInstanceInfo(),
		Maintenance:          false,
		MaintenanceReason:    "",
		PowerState:           "",
		TargetPowerState:     "",
		ProvisionState:       "",
		TargetProvisionState: "",
		ConsoleEnabled:       false,
		CreatedAt:            "",
		UpdatedAt:            "",
		LastError:            "",
	}
}

// MakeBaremetalNode makes BaremetalNode
// nolint
func InterfaceToBaremetalNode(i interface{}) *BaremetalNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BaremetalNode{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		Name:                 common.InterfaceToString(m["name"]),
		DriverInfo:           InterfaceToDriverInfo(m["driver_info"]),
		BMProperties:         InterfaceToBaremetalProperties(m["bm_properties"]),
		InstanceUUID:         common.InterfaceToString(m["instance_uuid"]),
		InstanceInfo:         InterfaceToInstanceInfo(m["instance_info"]),
		Maintenance:          common.InterfaceToBool(m["maintenance"]),
		MaintenanceReason:    common.InterfaceToString(m["maintenance_reason"]),
		PowerState:           common.InterfaceToString(m["power_state"]),
		TargetPowerState:     common.InterfaceToString(m["target_power_state"]),
		ProvisionState:       common.InterfaceToString(m["provision_state"]),
		TargetProvisionState: common.InterfaceToString(m["target_provision_state"]),
		ConsoleEnabled:       common.InterfaceToBool(m["console_enabled"]),
		CreatedAt:            common.InterfaceToString(m["created_at"]),
		UpdatedAt:            common.InterfaceToString(m["updated_at"]),
		LastError:            common.InterfaceToString(m["last_error"]),
	}
}

// MakeBaremetalNodeSlice() makes a slice of BaremetalNode
// nolint
func MakeBaremetalNodeSlice() []*BaremetalNode {
	return []*BaremetalNode{}
}

// InterfaceToBaremetalNodeSlice() makes a slice of BaremetalNode
// nolint
func InterfaceToBaremetalNodeSlice(i interface{}) []*BaremetalNode {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BaremetalNode{}
	for _, item := range list {
		result = append(result, InterfaceToBaremetalNode(item))
	}
	return result
}
