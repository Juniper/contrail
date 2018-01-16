package models

// BGPAsAService

import "encoding/json"

// BGPAsAService
type BGPAsAService struct {
	DisplayName                      string               `json:"display_name,omitempty"`
	Perms2                           *PermType2           `json:"perms2,omitempty"`
	ParentType                       string               `json:"parent_type,omitempty"`
	FQName                           []string             `json:"fq_name,omitempty"`
	IDPerms                          *IdPermsType         `json:"id_perms,omitempty"`
	Annotations                      *KeyValuePairs       `json:"annotations,omitempty"`
	UUID                             string               `json:"uuid,omitempty"`
	ParentUUID                       string               `json:"parent_uuid,omitempty"`
	BgpaasSessionAttributes          string               `json:"bgpaas_session_attributes,omitempty"`
	BgpaasSuppressRouteAdvertisement bool                 `json:"bgpaas_suppress_route_advertisement,omitempty"`
	BgpaasIpv4MappedIpv6Nexthop      bool                 `json:"bgpaas_ipv4_mapped_ipv6_nexthop,omitempty"`
	BgpaasIPAddress                  IpAddressType        `json:"bgpaas_ip_address,omitempty"`
	AutonomousSystem                 AutonomousSystemType `json:"autonomous_system,omitempty"`
	BgpaasShared                     bool                 `json:"bgpaas_shared,omitempty"`

	VirtualMachineInterfaceRefs []*BGPAsAServiceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	ServiceHealthCheckRefs      []*BGPAsAServiceServiceHealthCheckRef      `json:"service_health_check_refs,omitempty"`
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
		IDPerms:                          MakeIdPermsType(),
		Annotations:                      MakeKeyValuePairs(),
		UUID:                             "",
		BgpaasIPAddress:                  MakeIpAddressType(),
		AutonomousSystem:                 MakeAutonomousSystemType(),
		ParentUUID:                       "",
		BgpaasSessionAttributes:          "",
		BgpaasSuppressRouteAdvertisement: false,
		BgpaasIpv4MappedIpv6Nexthop:      false,
		BgpaasShared:                     false,
		FQName:                           []string{},
		DisplayName:                      "",
		Perms2:                           MakePermType2(),
		ParentType:                       "",
	}
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}
