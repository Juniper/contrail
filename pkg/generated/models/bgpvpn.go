package models

// BGPVPN

import "encoding/json"

// BGPVPN
type BGPVPN struct {
	BGPVPNType            VpnType          `json:"bgpvpn_type,omitempty"`
	UUID                  string           `json:"uuid,omitempty"`
	ParentType            string           `json:"parent_type,omitempty"`
	FQName                []string         `json:"fq_name,omitempty"`
	DisplayName           string           `json:"display_name,omitempty"`
	ImportRouteTargetList *RouteTargetList `json:"import_route_target_list,omitempty"`
	ExportRouteTargetList *RouteTargetList `json:"export_route_target_list,omitempty"`
	Perms2                *PermType2       `json:"perms2,omitempty"`
	ParentUUID            string           `json:"parent_uuid,omitempty"`
	IDPerms               *IdPermsType     `json:"id_perms,omitempty"`
	RouteTargetList       *RouteTargetList `json:"route_target_list,omitempty"`
	Annotations           *KeyValuePairs   `json:"annotations,omitempty"`
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
		UUID:                  "",
		ParentType:            "",
		FQName:                []string{},
		DisplayName:           "",
		ImportRouteTargetList: MakeRouteTargetList(),
		ExportRouteTargetList: MakeRouteTargetList(),
		Perms2:                MakePermType2(),
		ParentUUID:            "",
		IDPerms:               MakeIdPermsType(),
		RouteTargetList:       MakeRouteTargetList(),
		Annotations:           MakeKeyValuePairs(),
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}
