package models

// BGPAsAService

import "encoding/json"

// BGPAsAService
type BGPAsAService struct {
	AutonomousSystem                 AutonomousSystemType `json:"autonomous_system"`
	BgpaasShared                     bool                 `json:"bgpaas_shared"`
	ParentUUID                       string               `json:"parent_uuid"`
	IDPerms                          *IdPermsType         `json:"id_perms"`
	BgpaasSessionAttributes          string               `json:"bgpaas_session_attributes"`
	BgpaasSuppressRouteAdvertisement bool                 `json:"bgpaas_suppress_route_advertisement"`
	BgpaasIpv4MappedIpv6Nexthop      bool                 `json:"bgpaas_ipv4_mapped_ipv6_nexthop"`
	UUID                             string               `json:"uuid"`
	ParentType                       string               `json:"parent_type"`
	FQName                           []string             `json:"fq_name"`
	Annotations                      *KeyValuePairs       `json:"annotations"`
	BgpaasIPAddress                  IpAddressType        `json:"bgpaas_ip_address"`
	DisplayName                      string               `json:"display_name"`
	Perms2                           *PermType2           `json:"perms2"`

	ServiceHealthCheckRefs      []*BGPAsAServiceServiceHealthCheckRef      `json:"service_health_check_refs"`
	VirtualMachineInterfaceRefs []*BGPAsAServiceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
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
		UUID:                             "",
		ParentType:                       "",
		FQName:                           []string{},
		Annotations:                      MakeKeyValuePairs(),
		BgpaasSessionAttributes:          "",
		BgpaasSuppressRouteAdvertisement: false,
		BgpaasIpv4MappedIpv6Nexthop:      false,
		BgpaasIPAddress:                  MakeIpAddressType(),
		DisplayName:                      "",
		Perms2:                           MakePermType2(),
		AutonomousSystem:                 MakeAutonomousSystemType(),
		BgpaasShared:                     false,
		ParentUUID:                       "",
		IDPerms:                          MakeIdPermsType(),
	}
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}
