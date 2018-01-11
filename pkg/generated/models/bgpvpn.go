package models

// BGPVPN

import "encoding/json"

// BGPVPN
type BGPVPN struct {
	UUID                  string           `json:"uuid"`
	ParentType            string           `json:"parent_type"`
	FQName                []string         `json:"fq_name"`
	IDPerms               *IdPermsType     `json:"id_perms"`
	DisplayName           string           `json:"display_name"`
	RouteTargetList       *RouteTargetList `json:"route_target_list"`
	ImportRouteTargetList *RouteTargetList `json:"import_route_target_list"`
	ExportRouteTargetList *RouteTargetList `json:"export_route_target_list"`
	Perms2                *PermType2       `json:"perms2"`
	BGPVPNType            VpnType          `json:"bgpvpn_type"`
	ParentUUID            string           `json:"parent_uuid"`
	Annotations           *KeyValuePairs   `json:"annotations"`
}

// String returns json representation of the object
func (model *BGPVPN) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPVPN makes BGPVPN
func MakeBGPVPN() *BGPVPN {
	return &BGPVPN{
		//TODO(nati): Apply default
		ExportRouteTargetList: MakeRouteTargetList(),
		UUID:                  "",
		ParentType:            "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		RouteTargetList:       MakeRouteTargetList(),
		ImportRouteTargetList: MakeRouteTargetList(),
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		BGPVPNType:            MakeVpnType(),
		ParentUUID:            "",
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}
