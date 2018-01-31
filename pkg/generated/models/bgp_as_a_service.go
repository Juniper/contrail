package models

// BGPAsAService

import "encoding/json"

// BGPAsAService
type BGPAsAService struct {
	Perms2                           *PermType2           `json:"perms2,omitempty"`
	BgpaasSuppressRouteAdvertisement bool                 `json:"bgpaas_suppress_route_advertisement"`
	ParentUUID                       string               `json:"parent_uuid,omitempty"`
	DisplayName                      string               `json:"display_name,omitempty"`
	AutonomousSystem                 AutonomousSystemType `json:"autonomous_system,omitempty"`
	ParentType                       string               `json:"parent_type,omitempty"`
	UUID                             string               `json:"uuid,omitempty"`
	BgpaasSessionAttributes          string               `json:"bgpaas_session_attributes,omitempty"`
	FQName                           []string             `json:"fq_name,omitempty"`
	IDPerms                          *IdPermsType         `json:"id_perms,omitempty"`
	Annotations                      *KeyValuePairs       `json:"annotations,omitempty"`
	BgpaasShared                     bool                 `json:"bgpaas_shared"`
	BgpaasIpv4MappedIpv6Nexthop      bool                 `json:"bgpaas_ipv4_mapped_ipv6_nexthop"`
	BgpaasIPAddress                  IpAddressType        `json:"bgpaas_ip_address,omitempty"`

	VirtualMachineInterfaceRefs []*BGPAsAServiceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	ServiceHealthCheckRefs      []*BGPAsAServiceServiceHealthCheckRef      `json:"service_health_check_refs,omitempty"`
}

// BGPAsAServiceServiceHealthCheckRef references each other
type BGPAsAServiceServiceHealthCheckRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// BGPAsAServiceVirtualMachineInterfaceRef references each other
type BGPAsAServiceVirtualMachineInterfaceRef struct {
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
		BgpaasSessionAttributes:     "",
		Annotations:                 MakeKeyValuePairs(),
		BgpaasShared:                false,
		BgpaasIpv4MappedIpv6Nexthop: false,
		BgpaasIPAddress:             MakeIpAddressType(),
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		BgpaasSuppressRouteAdvertisement: false,
		ParentUUID:                       "",
		DisplayName:                      "",
		Perms2:                           MakePermType2(),
		AutonomousSystem:                 MakeAutonomousSystemType(),
		ParentType:                       "",
		UUID:                             "",
	}
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}
