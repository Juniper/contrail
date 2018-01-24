package models

// BGPAsAService

import "encoding/json"

// BGPAsAService
type BGPAsAService struct {
	UUID                             string               `json:"uuid,omitempty"`
	FQName                           []string             `json:"fq_name,omitempty"`
	DisplayName                      string               `json:"display_name,omitempty"`
	BgpaasIpv4MappedIpv6Nexthop      bool                 `json:"bgpaas_ipv4_mapped_ipv6_nexthop"`
	Annotations                      *KeyValuePairs       `json:"annotations,omitempty"`
	BgpaasIPAddress                  IpAddressType        `json:"bgpaas_ip_address,omitempty"`
	IDPerms                          *IdPermsType         `json:"id_perms,omitempty"`
	BgpaasSuppressRouteAdvertisement bool                 `json:"bgpaas_suppress_route_advertisement"`
	BgpaasSessionAttributes          string               `json:"bgpaas_session_attributes,omitempty"`
	AutonomousSystem                 AutonomousSystemType `json:"autonomous_system,omitempty"`
	Perms2                           *PermType2           `json:"perms2,omitempty"`
	ParentUUID                       string               `json:"parent_uuid,omitempty"`
	ParentType                       string               `json:"parent_type,omitempty"`
	BgpaasShared                     bool                 `json:"bgpaas_shared"`

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
		BgpaasIpv4MappedIpv6Nexthop: false,
		UUID:                             "",
		FQName:                           []string{},
		DisplayName:                      "",
		Annotations:                      MakeKeyValuePairs(),
		BgpaasSuppressRouteAdvertisement: false,
		BgpaasIPAddress:                  MakeIpAddressType(),
		IDPerms:                          MakeIdPermsType(),
		ParentType:                       "",
		BgpaasShared:                     false,
		BgpaasSessionAttributes:          "",
		AutonomousSystem:                 MakeAutonomousSystemType(),
		Perms2:                           MakePermType2(),
		ParentUUID:                       "",
	}
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}
