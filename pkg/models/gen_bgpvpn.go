package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeBGPVPN makes BGPVPN
// nolint
func MakeBGPVPN() *BGPVPN {
	return &BGPVPN{
		//TODO(nati): Apply default
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		RouteTargetList:       MakeRouteTargetList(),
		ImportRouteTargetList: MakeRouteTargetList(),
		ExportRouteTargetList: MakeRouteTargetList(),
		BGPVPNType:            "",
	}
}

// MakeBGPVPN makes BGPVPN
// nolint
func InterfaceToBGPVPN(i interface{}) *BGPVPN {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BGPVPN{
		//TODO(nati): Apply default
		UUID:                  common.InterfaceToString(m["uuid"]),
		ParentUUID:            common.InterfaceToString(m["parent_uuid"]),
		ParentType:            common.InterfaceToString(m["parent_type"]),
		FQName:                common.InterfaceToStringList(m["fq_name"]),
		IDPerms:               InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:           common.InterfaceToString(m["display_name"]),
		Annotations:           InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                InterfaceToPermType2(m["perms2"]),
		RouteTargetList:       InterfaceToRouteTargetList(m["route_target_list"]),
		ImportRouteTargetList: InterfaceToRouteTargetList(m["import_route_target_list"]),
		ExportRouteTargetList: InterfaceToRouteTargetList(m["export_route_target_list"]),
		BGPVPNType:            common.InterfaceToString(m["bgpvpn_type"]),
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
// nolint
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}

// InterfaceToBGPVPNSlice() makes a slice of BGPVPN
// nolint
func InterfaceToBGPVPNSlice(i interface{}) []*BGPVPN {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BGPVPN{}
	for _, item := range list {
		result = append(result, InterfaceToBGPVPN(item))
	}
	return result
}
