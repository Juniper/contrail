package models

// BGPVPN

import "encoding/json"

// BGPVPN
type BGPVPN struct {
	Perms2                *PermType2       `json:"perms2,omitempty"`
	UUID                  string           `json:"uuid,omitempty"`
	ParentUUID            string           `json:"parent_uuid,omitempty"`
	ParentType            string           `json:"parent_type,omitempty"`
	ExportRouteTargetList *RouteTargetList `json:"export_route_target_list,omitempty"`
	Annotations           *KeyValuePairs   `json:"annotations,omitempty"`
	BGPVPNType            VpnType          `json:"bgpvpn_type,omitempty"`
	IDPerms               *IdPermsType     `json:"id_perms,omitempty"`
	DisplayName           string           `json:"display_name,omitempty"`
	FQName                []string         `json:"fq_name,omitempty"`
	RouteTargetList       *RouteTargetList `json:"route_target_list,omitempty"`
	ImportRouteTargetList *RouteTargetList `json:"import_route_target_list,omitempty"`
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
		ImportRouteTargetList: MakeRouteTargetList(),
		BGPVPNType:            MakeVpnType(),
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		FQName:                []string{},
		RouteTargetList:       MakeRouteTargetList(),
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
		ExportRouteTargetList: MakeRouteTargetList(),
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}
