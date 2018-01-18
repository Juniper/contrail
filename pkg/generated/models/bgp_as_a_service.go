package models

// BGPAsAService

import "encoding/json"

// BGPAsAService
type BGPAsAService struct {
	BgpaasSessionAttributes          string               `json:"bgpaas_session_attributes,omitempty"`
	Annotations                      *KeyValuePairs       `json:"annotations,omitempty"`
	BgpaasIpv4MappedIpv6Nexthop      bool                 `json:"bgpaas_ipv4_mapped_ipv6_nexthop"`
	FQName                           []string             `json:"fq_name,omitempty"`
	ParentType                       string               `json:"parent_type,omitempty"`
	BgpaasSuppressRouteAdvertisement bool                 `json:"bgpaas_suppress_route_advertisement"`
	BgpaasIPAddress                  IpAddressType        `json:"bgpaas_ip_address,omitempty"`
	AutonomousSystem                 AutonomousSystemType `json:"autonomous_system,omitempty"`
	IDPerms                          *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName                      string               `json:"display_name,omitempty"`
	BgpaasShared                     bool                 `json:"bgpaas_shared"`
	Perms2                           *PermType2           `json:"perms2,omitempty"`
	UUID                             string               `json:"uuid,omitempty"`
	ParentUUID                       string               `json:"parent_uuid,omitempty"`

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
		DisplayName:                      "",
		BgpaasSuppressRouteAdvertisement: false,
		BgpaasIPAddress:                  MakeIpAddressType(),
		AutonomousSystem:                 MakeAutonomousSystemType(),
		ParentUUID:                       "",
		BgpaasShared:                     false,
		Perms2:                           MakePermType2(),
		UUID:                             "",
		BgpaasSessionAttributes:     "",
		Annotations:                 MakeKeyValuePairs(),
		BgpaasIpv4MappedIpv6Nexthop: false,
		FQName:     []string{},
		ParentType: "",
	}
}

// MakeBGPAsAServiceSlice() makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
	return []*BGPAsAService{}
}
