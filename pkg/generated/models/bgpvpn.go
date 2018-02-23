package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeBGPVPN makes BGPVPN
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
func InterfaceToBGPVPN(i interface{}) *BGPVPN {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &BGPVPN{
		//TODO(nati): Apply default
		UUID:                  schema.InterfaceToString(m["uuid"]),
		ParentUUID:            schema.InterfaceToString(m["parent_uuid"]),
		ParentType:            schema.InterfaceToString(m["parent_type"]),
		FQName:                schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:               InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:           schema.InterfaceToString(m["display_name"]),
		Annotations:           InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                InterfaceToPermType2(m["perms2"]),
		RouteTargetList:       InterfaceToRouteTargetList(m["route_target_list"]),
		ImportRouteTargetList: InterfaceToRouteTargetList(m["import_route_target_list"]),
		ExportRouteTargetList: InterfaceToRouteTargetList(m["export_route_target_list"]),
		BGPVPNType:            schema.InterfaceToString(m["bgpvpn_type"]),
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}

// InterfaceToBGPVPNSlice() makes a slice of BGPVPN
func InterfaceToBGPVPNSlice(i interface{}) []*BGPVPN {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*BGPVPN{}
	for _, item := range list {
		result = append(result, InterfaceToBGPVPN(item))
	}
	return result
}
