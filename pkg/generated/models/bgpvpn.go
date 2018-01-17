package models

// BGPVPN

import "encoding/json"

// BGPVPN
type BGPVPN struct {
	Annotations           *KeyValuePairs   `json:"annotations,omitempty"`
	Perms2                *PermType2       `json:"perms2,omitempty"`
	RouteTargetList       *RouteTargetList `json:"route_target_list,omitempty"`
	ImportRouteTargetList *RouteTargetList `json:"import_route_target_list,omitempty"`
	BGPVPNType            VpnType          `json:"bgpvpn_type,omitempty"`
	DisplayName           string           `json:"display_name,omitempty"`
	ParentUUID            string           `json:"parent_uuid,omitempty"`
	ParentType            string           `json:"parent_type,omitempty"`
	ExportRouteTargetList *RouteTargetList `json:"export_route_target_list,omitempty"`
	FQName                []string         `json:"fq_name,omitempty"`
	IDPerms               *IdPermsType     `json:"id_perms,omitempty"`
	UUID                  string           `json:"uuid,omitempty"`
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
		RouteTargetList:       MakeRouteTargetList(),
		ImportRouteTargetList: MakeRouteTargetList(),
		BGPVPNType:            MakeVpnType(),
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		ExportRouteTargetList: MakeRouteTargetList(),
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
	}
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}
