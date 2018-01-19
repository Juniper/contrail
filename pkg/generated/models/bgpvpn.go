package models

// BGPVPN

import "encoding/json"

// BGPVPN
type BGPVPN struct {
	RouteTargetList       *RouteTargetList `json:"route_target_list,omitempty"`
	ImportRouteTargetList *RouteTargetList `json:"import_route_target_list,omitempty"`
	BGPVPNType            VpnType          `json:"bgpvpn_type,omitempty"`
	UUID                  string           `json:"uuid,omitempty"`
	FQName                []string         `json:"fq_name,omitempty"`
	IDPerms               *IdPermsType     `json:"id_perms,omitempty"`
	Annotations           *KeyValuePairs   `json:"annotations,omitempty"`
	ExportRouteTargetList *RouteTargetList `json:"export_route_target_list,omitempty"`
	Perms2                *PermType2       `json:"perms2,omitempty"`
	ParentUUID            string           `json:"parent_uuid,omitempty"`
	ParentType            string           `json:"parent_type,omitempty"`
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
		UUID:                  "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		Annotations:           MakeKeyValuePairs(),
		RouteTargetList:       MakeRouteTargetList(),
		ImportRouteTargetList: MakeRouteTargetList(),
		BGPVPNType:            MakeVpnType(),
		ParentType:            "",
		DisplayName:           "",
		ExportRouteTargetList: MakeRouteTargetList(),
		Perms2:                MakePermType2(),
		ParentUUID:            "",
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}
