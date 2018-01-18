package models

// BGPVPN

import "encoding/json"

// BGPVPN
type BGPVPN struct {
	ParentUUID            string           `json:"parent_uuid,omitempty"`
	ParentType            string           `json:"parent_type,omitempty"`
	BGPVPNType            VpnType          `json:"bgpvpn_type,omitempty"`
	IDPerms               *IdPermsType     `json:"id_perms,omitempty"`
	Annotations           *KeyValuePairs   `json:"annotations,omitempty"`
	Perms2                *PermType2       `json:"perms2,omitempty"`
	UUID                  string           `json:"uuid,omitempty"`
	FQName                []string         `json:"fq_name,omitempty"`
	RouteTargetList       *RouteTargetList `json:"route_target_list,omitempty"`
	ImportRouteTargetList *RouteTargetList `json:"import_route_target_list,omitempty"`
	ExportRouteTargetList *RouteTargetList `json:"export_route_target_list,omitempty"`
	DisplayName           string           `json:"display_name,omitempty"`
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
		BGPVPNType:            MakeVpnType(),
		IDPerms:               MakeIdPermsType(),
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		ParentUUID:            "",
		ParentType:            "",
		RouteTargetList:       MakeRouteTargetList(),
		ImportRouteTargetList: MakeRouteTargetList(),
		ExportRouteTargetList: MakeRouteTargetList(),
		DisplayName:           "",
		UUID:                  "",
		FQName:                []string{},
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}
