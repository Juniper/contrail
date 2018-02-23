package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeBridgeDomain makes BridgeDomain
func MakeBridgeDomain() *BridgeDomain {
	return &BridgeDomain{
		//TODO(nati): Apply default
		UUID:               "",
		ParentUUID:         "",
		ParentType:         "",
		FQName:             []string{},
		IDPerms:            MakeIdPermsType(),
		DisplayName:        "",
		Annotations:        MakeKeyValuePairs(),
		Perms2:             MakePermType2(),
		MacAgingTime:       0,
		Isid:               0,
		MacLearningEnabled: false,
		MacMoveControl:     MakeMACMoveLimitControlType(),
		MacLimitControl:    MakeMACLimitControlType(),
	}
}

// MakeBridgeDomain makes BridgeDomain
func InterfaceToBridgeDomain(i interface{}) *BridgeDomain {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BridgeDomain{
		//TODO(nati): Apply default
		UUID:               schema.InterfaceToString(m["uuid"]),
		ParentUUID:         schema.InterfaceToString(m["parent_uuid"]),
		ParentType:         schema.InterfaceToString(m["parent_type"]),
		FQName:             schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:            InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:        schema.InterfaceToString(m["display_name"]),
		Annotations:        InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:             InterfaceToPermType2(m["perms2"]),
		MacAgingTime:       schema.InterfaceToInt64(m["mac_aging_time"]),
		Isid:               schema.InterfaceToInt64(m["isid"]),
		MacLearningEnabled: schema.InterfaceToBool(m["mac_learning_enabled"]),
		MacMoveControl:     InterfaceToMACMoveLimitControlType(m["mac_move_control"]),
		MacLimitControl:    InterfaceToMACLimitControlType(m["mac_limit_control"]),
	}
}

// MakeBridgeDomainSlice() makes a slice of BridgeDomain
func MakeBridgeDomainSlice() []*BridgeDomain {
	return []*BridgeDomain{}
}

// InterfaceToBridgeDomainSlice() makes a slice of BridgeDomain
func InterfaceToBridgeDomainSlice(i interface{}) []*BridgeDomain {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BridgeDomain{}
	for _, item := range list {
		result = append(result, InterfaceToBridgeDomain(item))
	}
	return result
}
