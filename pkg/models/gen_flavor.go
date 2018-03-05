package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFlavor makes Flavor
// nolint
func MakeFlavor() *Flavor {
	return &Flavor{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
		Name:                 "",
		Disk:                 0,
		Vcpus:                0,
		RAM:                  0,
		ID:                   "",
		Property:             "",
		RXTXFactor:           0,
		Swap:                 0,
		IsPublic:             false,
		Ephemeral:            0,
		Links:                MakeOpenStackLink(),
	}
}

// MakeFlavor makes Flavor
// nolint
func InterfaceToFlavor(i interface{}) *Flavor {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Flavor{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),
		Name:                 common.InterfaceToString(m["name"]),
		Disk:                 common.InterfaceToInt64(m["disk"]),
		Vcpus:                common.InterfaceToInt64(m["vcpus"]),
		RAM:                  common.InterfaceToInt64(m["ram"]),
		ID:                   common.InterfaceToString(m["id"]),
		Property:             common.InterfaceToString(m["property"]),
		RXTXFactor:           common.InterfaceToInt64(m["rxtx_factor"]),
		Swap:                 common.InterfaceToInt64(m["swap"]),
		IsPublic:             common.InterfaceToBool(m["is_public"]),
		Ephemeral:            common.InterfaceToInt64(m["ephemeral"]),
		Links:                InterfaceToOpenStackLink(m["links"]),
	}
}

// MakeFlavorSlice() makes a slice of Flavor
// nolint
func MakeFlavorSlice() []*Flavor {
	return []*Flavor{}
}

// InterfaceToFlavorSlice() makes a slice of Flavor
// nolint
func InterfaceToFlavorSlice(i interface{}) []*Flavor {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Flavor{}
	for _, item := range list {
		result = append(result, InterfaceToFlavor(item))
	}
	return result
}
