package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeKeypair makes Keypair
// nolint
func MakeKeypair() *Keypair {
	return &Keypair{
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
		PublicKey:            "",
	}
}

// MakeKeypair makes Keypair
// nolint
func InterfaceToKeypair(i interface{}) *Keypair {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Keypair{
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
		PublicKey:            common.InterfaceToString(m["public_key"]),
	}
}

// MakeKeypairSlice() makes a slice of Keypair
// nolint
func MakeKeypairSlice() []*Keypair {
	return []*Keypair{}
}

// InterfaceToKeypairSlice() makes a slice of Keypair
// nolint
func InterfaceToKeypairSlice(i interface{}) []*Keypair {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Keypair{}
	for _, item := range list {
		result = append(result, InterfaceToKeypair(item))
	}
	return result
}
