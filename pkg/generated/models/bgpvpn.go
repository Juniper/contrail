package models

// BGPVPN

import "encoding/json"

// BGPVPN
type BGPVPN struct {
	UUID                  string           `json:"uuid"`
	ImportRouteTargetList *RouteTargetList `json:"import_route_target_list"`
	ExportRouteTargetList *RouteTargetList `json:"export_route_target_list"`
	ParentUUID            string           `json:"parent_uuid"`
	FQName                []string         `json:"fq_name"`
	Annotations           *KeyValuePairs   `json:"annotations"`
	Perms2                *PermType2       `json:"perms2"`
	RouteTargetList       *RouteTargetList `json:"route_target_list"`
	BGPVPNType            VpnType          `json:"bgpvpn_type"`
	ParentType            string           `json:"parent_type"`
	IDPerms               *IdPermsType     `json:"id_perms"`
	DisplayName           string           `json:"display_name"`
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
		ExportRouteTargetList: MakeRouteTargetList(),
		ParentUUID:            "",
		FQName:                []string{},
		Annotations:           MakeKeyValuePairs(),
		UUID:                  "",
		RouteTargetList:       MakeRouteTargetList(),
		BGPVPNType:            MakeVpnType(),
		ParentType:            "",
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Perms2:                MakePermType2(),
	}
}

// InterfaceToBGPVPN makes BGPVPN from interface
func InterfaceToBGPVPN(iData interface{}) *BGPVPN {
	data := iData.(map[string]interface{})
	return &BGPVPN{
		RouteTargetList: InterfaceToRouteTargetList(data["route_target_list"]),

		//{"description":"List of route targets that are used as both import and export for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		BGPVPNType: InterfaceToVpnType(data["bgpvpn_type"]),

		//{"description":"BGP VPN type selection between IP VPN (l3) and Ethernet VPN (l2) (default: l3).","default":"l3","type":"string","enum":["l2","l3"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ImportRouteTargetList: InterfaceToRouteTargetList(data["import_route_target_list"]),

		//{"description":"List of route targets that are used as import for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		ExportRouteTargetList: InterfaceToRouteTargetList(data["export_route_target_list"]),

		//{"description":"List of route targets that are used as export for this virtual network.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToBGPVPNSlice makes a slice of BGPVPN from interface
func InterfaceToBGPVPNSlice(data interface{}) []*BGPVPN {
	list := data.([]interface{})
	result := MakeBGPVPNSlice()
	for _, item := range list {
		result = append(result, InterfaceToBGPVPN(item))
	}
	return result
}

// MakeBGPVPNSlice() makes a slice of BGPVPN
func MakeBGPVPNSlice() []*BGPVPN {
	return []*BGPVPN{}
}
