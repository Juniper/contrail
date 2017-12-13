package models

// BGPAsAService

import "encoding/json"

// BGPAsAService
type BGPAsAService struct {
	Annotations                      *KeyValuePairs       `json:"annotations"`
	BgpaasShared                     bool                 `json:"bgpaas_shared"`
	BgpaasSessionAttributes          string               `json:"bgpaas_session_attributes"`
	BgpaasIpv4MappedIpv6Nexthop      bool                 `json:"bgpaas_ipv4_mapped_ipv6_nexthop"`
	DisplayName                      string               `json:"display_name"`
	BgpaasSuppressRouteAdvertisement bool                 `json:"bgpaas_suppress_route_advertisement"`
	AutonomousSystem                 AutonomousSystemType `json:"autonomous_system"`
	ParentUUID                       string               `json:"parent_uuid"`
	IDPerms                          *IdPermsType         `json:"id_perms"`
	UUID                             string               `json:"uuid"`
	BgpaasIPAddress                  IpAddressType        `json:"bgpaas_ip_address"`
	ParentType                       string               `json:"parent_type"`
	FQName                           []string             `json:"fq_name"`
	Perms2                           *PermType2           `json:"perms2"`

	VirtualMachineInterfaceRefs []*BGPAsAServiceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
	ServiceHealthCheckRefs      []*BGPAsAServiceServiceHealthCheckRef      `json:"service_health_check_refs"`
}

// BGPAsAServiceVirtualMachineInterfaceRef references each other
type BGPAsAServiceVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// BGPAsAServiceServiceHealthCheckRef references each other
type BGPAsAServiceServiceHealthCheckRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *BGPAsAService) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBGPAsAService makes BGPAsAService
func MakeBGPAsAService() *BGPAsAService {
	return &BGPAsAService{
		//TODO(nati): Apply default
		Annotations:                      MakeKeyValuePairs(),
		BgpaasShared:                     false,
		BgpaasSessionAttributes:          "",
		BgpaasIpv4MappedIpv6Nexthop:      false,
		DisplayName:                      "",
		BgpaasSuppressRouteAdvertisement: false,
		AutonomousSystem:                 MakeAutonomousSystemType(),
		ParentUUID:                       "",
		IDPerms:                          MakeIdPermsType(),
		UUID:                             "",
		BgpaasIPAddress:                  MakeIpAddressType(),
		ParentType:                       "",
		FQName:                           []string{},
		Perms2:                           MakePermType2(),
	}
}

// InterfaceToBGPAsAService makes BGPAsAService from interface
func InterfaceToBGPAsAService(iData interface{}) *BGPAsAService {
	data := iData.(map[string]interface{})
	return &BGPAsAService{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		BgpaasSessionAttributes: data["bgpaas_session_attributes"].(string),

		//{"description":"BGP peering session attributes.","type":"string"}
		BgpaasIpv4MappedIpv6Nexthop: data["bgpaas_ipv4_mapped_ipv6_nexthop"].(bool),

		//{"description":"True when client bgp implementation expects to receive a ipv4-mapped ipv6 address (as opposed to regular ipv6 address) as the bgp nexthop for ipv6 routes.","type":"boolean"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		BgpaasShared: data["bgpaas_shared"].(bool),

		//{"description":"True if only one BGP router needs to be created. Otherwise, one BGP router is created for each VMI","default":false,"type":"boolean"}
		AutonomousSystem: InterfaceToAutonomousSystemType(data["autonomous_system"]),

		//{"description":"16 bit BGP Autonomous System number for the cluster.","type":"integer","minimum":1,"maximum":65534}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		BgpaasSuppressRouteAdvertisement: data["bgpaas_suppress_route_advertisement"].(bool),

		//{"description":"True when server should not advertise any routes to the client i.e. the client has static routes (typically a default) configured.","type":"boolean"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		BgpaasIPAddress: InterfaceToIpAddressType(data["bgpaas_ip_address"]),

		//{"description":"Ip address of the BGP peer.","type":"string"}

	}
}

// InterfaceToBGPAsAServiceSlice makes a slice of BGPAsAService from interface
func InterfaceToBGPAsAServiceSlice(data interface{}) []*BGPAsAService {
	list := data.([]interface{})
	result := MakeBGPAsAServiceSlice()
	for _, item := range list {
		result = append(result, InterfaceToBGPAsAService(item))
	}
	return result
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}
